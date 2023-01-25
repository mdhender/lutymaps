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

// Package main implements the mapping engine for luty.
package main

import (
	"fmt"
	"github.com/mdhender/lutymaps/store/jsdb"
	"github.com/mdhender/lutymaps/store/mem"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func run() error {
	jstore, err := jsdb.New("galaxy-001.json")
	if err != nil {
		return fmt.Errorf("luty: %w", err)
	}

	mstore, err := mem.AdaptJSDBToStore(jstore)
	if err != nil {
		return fmt.Errorf("luty: %w", err)
	}
	if mstore == nil {
		panic("assert(mstore != nil)")
	}

	jstore, err = mem.AdaptStoreToJSDB(mstore)
	if err != nil {
		return fmt.Errorf("luty: %w", err)
	}
	err = jstore.Save("galaxy-002.json")
	if err != nil {
		return fmt.Errorf("luty: %w", err)
	}

	return nil
}