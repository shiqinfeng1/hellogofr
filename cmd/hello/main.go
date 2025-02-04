// Copyright @2025-2028 <SieYuan> . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shiqinfeng1/hellogofr/internal/hello/server"

	_ "go.uber.org/automaxprocs"
)

func main() {
	server.New().Run()
}
