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

package jsdb

import (
	"encoding/json"
	"fmt"
	"os"
)

// AccountStore implements a flat file data store using JSON.
type AccountStore struct {
	Accounts []*Account `json:"accounts"`
}

// Account details
type Account struct {
	Id     string   `json:"id"`
	UserId string   `json:"user-id"`
	Secret string   `json:"secret"`
	Roles  []string `json:"roles"`
}

// Load loads the store from the path.
func (s *AccountStore) Load(path string) error {
	s.Accounts = nil
	buf, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("jsdb: %w", err)
	}
	err = json.Unmarshal(buf, &s)
	if err != nil {
		return fmt.Errorf("jsdb: %w", err)
	}
	return nil
}

// Save writes the store to the path.
func (s *AccountStore) Save(path string) error {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("jsdb: %w", err)
	}
	err = os.WriteFile(path, buf, 0666)
	if err != nil {
		return fmt.Errorf("jsdb: %w", err)
	}
	return nil
}
