/*
 * lutymaps - a mapping engine for luty
 *
 * Copyright (c) 2023 2023 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Package auth implements examples for authentication and authorization
package auth

// DO NOT USE - this is an example only!

// DoNotUse provides an example of how to implement these interfaces.
type DoNotUse struct{}

// Authenticate implements the server.Authentication interface.
func (dnu DoNotUse) Authenticate(id, secret string) (string, bool) {
	if id == "whiskey" && secret == "tango.foxtrot" {
		return "00112233-4455-6677-8899-aabbccddeeff", true
	}
	return "", false
}

// Authorize implements the server.Authorization interface.
func (dnu DoNotUse) Authorize(id string) func(role string) bool {
	if id == "00112233-4455-6677-8899-aabbccddeeff" {
		return func(role string) bool {
			return role == "guest"
		}
	}
	return func(_ string) bool {
		return false
	}
}
