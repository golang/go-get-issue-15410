// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", git)
}

const out = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="HOST/pkg git SCHEME://HOST/">
</head>
<body>
If you must know, this is
<a href="https://github.com/rsc/go-get-issue-15410">https://github.com/rsc/go-get-issue-15410</a>...
</body>
</html>
`

func slash(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Replace(strings.Replace(out, "SCHEME", r.URL.Scheme, -1), "HOST", r.Host, -1)))
}

func urlMustParse(text string) *url.URL {
	u, err := url.Parse(text)
	if err != nil {
		panic(err)
	}
	return u
}

func git(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		time.Sleep(20 * time.Second) // go get will time out
		http.Error(w, "tls not allowed", 500)
		return
	}
	if r.URL.Path == "/" || r.URL.Path == "/pkg" || r.URL.Path == "/pkg/p" {
		slash(w, r)
		return
	}
	ctx := appengine.NewContext(r)
	reverse := httputil.NewSingleHostReverseProxy(urlMustParse("https://github.com/rsc/go-get-issue-15410"))
	reverse.Transport = &urlfetch.Transport{Context: ctx}

	reverse.ServeHTTP(w, r)
}
