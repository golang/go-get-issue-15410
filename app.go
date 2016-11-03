// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", slash)
}

const out = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="HOST/pkg git https://github.com/rsc/go-get-issue-15410">
</head>
<body>
If you must know, this is
<a href="https://github.com/rsc/go-get-issue-15410">https://github.com/rsc/go-get-issue-15410</a>...
</body>
</html>
`

func slash(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Replace(out, "HOST", r.Host, -1)))
}
