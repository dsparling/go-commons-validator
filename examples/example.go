// Copyright 2013 Doug Sparling. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/dsparling/go-commons-validator/email"
)

func main() {
	// true
	fmt.Println(email.IsValid("test@example.com"))

	// false
	fmt.Println(email.IsValid("testexample.com"))
}
