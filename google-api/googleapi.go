// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package googleapi contains the common code shared by all Google API
// libraries.
package googleapi

import (
	"bytes"
	"fmt"
	"http"
	"io"
	"io/ioutil"
	"json"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"
	"url"
)

// ContentTyper is an interface for Readers which know (or would like
// to override) their Content-Type. If a media body doesn't implement
// ContentTyper, the type is sniffed from the content using
// http.DetectContentType.
type ContentTyper interface {
	ContentType() string
}

const Version = "0.5"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) String() string {
	return fmt.Sprintf("googleapi: Error %d: %s", e.Code, e.Message)
}

type errorReply struct {
	Error *Error `json:"error"`
}

func CheckResponse(res *http.Response) os.Error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(res.Body)
	if err == nil {
		jerr := new(errorReply)
		err = json.Unmarshal(slurp, jerr)
		if err == nil && jerr.Error != nil {
			return jerr.Error
		}
	}
	return fmt.Errorf("googleapi: got HTTP response code %d and error reading body: %v",
		res.StatusCode, err)
}

type MarshalStyle bool

var WithDataWrapper = MarshalStyle(true)
var WithoutDataWrapper = MarshalStyle(false)

func (wrap MarshalStyle) JSONReader(v interface{}) (io.Reader, os.Error) {
	buf := new(bytes.Buffer)
	if wrap {
		buf.Write([]byte(`{"data": `))
	}
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	if wrap {
		buf.Write([]byte(`}`))
	}
	return buf, nil
}

func getMediaType(media io.Reader) (io.Reader, string) {
	if typer, ok := media.(ContentTyper); ok {
		return media, typer.ContentType()
	}

	typ := "application/octet-stream"
	buf := make([]byte, 1024)
	n, err := media.Read(buf)
	buf = buf[:n]
	if err == nil {
		typ = http.DetectContentType(buf)
	}
	return io.MultiReader(bytes.NewBuffer(buf), media), typ
}

type Lengther interface {
	Len() int
}

// endingWithErrorReader from r until it returns an error.  If the
// final error from r is os.EOF and e is non-nil, e is used instead.
type endingWithErrorReader struct {
	r io.Reader
	e os.Error
}

func (er endingWithErrorReader) Read(p []byte) (n int, err os.Error) {
	n, err = er.r.Read(p)
	if err == os.EOF && er.e != nil {
		err = er.e
	}
	return
}

func getReaderSize(r io.Reader) (io.Reader, int64) {
	// Ideal case, the reader knows its own size.
	if lr, ok := r.(Lengther); ok {
		return r, int64(lr.Len())
	}

	// But maybe it's a seeker and we can seek to the end to find its size.
	if s, ok := r.(io.Seeker); ok {
		pos0, err := s.Seek(0, os.SEEK_CUR)
		if err == nil {
			posend, err := s.Seek(0, os.SEEK_END)
			if err == nil {
				_, err = s.Seek(pos0, os.SEEK_SET)
				if err == nil {
					return r, posend - pos0
				} else {
					// We moved it forward but can't restore it.
					// Seems unlikely, but can't really restore now.
					return endingWithErrorReader{strings.NewReader(""), err}, posend - pos0
				}
			}
		}
	}

	// Otherwise we have to make a copy to calculate how big the reader is.
	buf := new(bytes.Buffer)
	// TODO(bradfitz): put a cap on this copy? spill to disk after
	// a certain point?
	_, err := io.Copy(buf, r)
	return endingWithErrorReader{buf, err}, int64(buf.Len())
}

func typeHeader(contentType string) textproto.MIMEHeader {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", contentType)
	return h
}

// countingWriter counts the number of bytes it receives to write, but
// discards them.
type countingWriter struct {
	n *int64
}

func (w countingWriter) Write(p []byte) (int, os.Error) {
	*w.n += int64(len(p))
	return len(p), nil
}

// ConditionallyIncludeMedia does nothing if media is nil.
//
// bodyp is an in/out parameter.  It should initially point to the
// reader of the application/json (or whatever) payload to send in the
// API request.  It's updated to point to the multipart body reader.
//
// ctypep is an in/out parameter.  It should initially point to the
// content type of the bodyp, usually "application/json".  It's updated
// to the "multipart/related" content type, with random boundary.
//
// The return value is the content-length of the entire multpart body.
func ConditionallyIncludeMedia(media io.Reader, bodyp *io.Reader, ctypep *string) (totalContentLength int64, ok bool) {
	if media == nil {
		return
	}
	// Get the media type and size. The type check might return a
	// different reader instance, so do the size check first,
	// which looks at the specific type of the io.Reader.
	media, mediaSize := getReaderSize(media)
	media, mediaType := getMediaType(media)
	body, bodyType := *bodyp, *ctypep
	body, bodySize := getReaderSize(body)

	// Calculate how big the the multipart will be.
	{
		totalContentLength = bodySize + mediaSize
		mpw := multipart.NewWriter(countingWriter{&totalContentLength})
		mpw.CreatePart(typeHeader(bodyType))
		mpw.CreatePart(typeHeader(mediaType))
		mpw.Close()
	}

	pr, pw := io.Pipe()
	mpw := multipart.NewWriter(pw)
	*bodyp = pr
	*ctypep = "multipart/related; boundary=" + mpw.Boundary()
	go func() {
		defer pw.Close()
		defer mpw.Close()

		w, err := mpw.CreatePart(typeHeader(bodyType))
		if err != nil {
			return
		}
		_, err = io.Copy(w, body)
		if err != nil {
			return
		}

		w, err = mpw.CreatePart(typeHeader(mediaType))
		if err != nil {
			return
		}
		_, err = io.Copy(w, media)
		if err != nil {
			return
		}
	}()
	return totalContentLength, true
}

func ResolveRelative(basestr, relstr string) string {
	u, _ := url.Parse(basestr)
	rel, _ := url.Parse(relstr)
	u = u.ResolveReference(rel)
	us := u.String()
	us = strings.Replace(us, "%7B", "{", -1)
	us = strings.Replace(us, "%7D", "}", -1)
	return us
}
