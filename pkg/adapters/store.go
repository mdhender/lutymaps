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

package adapters

import (
	"github.com/mdhender/lutymaps/pkg/stores/jsdb"
	"github.com/mdhender/lutymaps/pkg/stores/mem"
	"sort"
)

// JSAccountsToMemAccounts converts JSON accounts to in-memory accounts
func JSAccountsToMemAccounts(js jsdb.AccountStore) (mem.Accounts, error) {
	accts := make(map[string]mem.Account)
	for _, acct := range js.Accounts {
		a := mem.Account{
			Id:           acct.Id,
			UserId:       acct.UserId,
			HashedSecret: acct.Secret,
			Roles:        make(map[string]bool),
		}
		for _, role := range acct.Roles {
			a.Roles[role] = true
		}
		accts[a.Id] = a
	}
	return accts, nil
}

// MemAccountsToJSAccounts converts in-memory accounts to JSON accounts
func MemAccountsToJSAccounts(ms mem.Accounts) (*jsdb.AccountStore, error) {
	js := &jsdb.AccountStore{}
	for _, acct := range ms {
		a := &jsdb.Account{
			Id:     acct.Id,
			UserId: acct.UserId,
			Secret: acct.HashedSecret,
		}
		for role, ok := range acct.Roles {
			if ok {
				a.Roles = append(a.Roles, role)
			}
		}
		sort.Strings(a.Roles)
		js.Accounts = append(js.Accounts, a)
	}
	sort.Slice(js.Accounts, func(i, j int) bool { return js.Accounts[i].Id < js.Accounts[j].Id })
	return js, nil
}

// JSDBToStore converts a JSDB store to an in-memory store.
func JSDBToStore(store *jsdb.Store) (*mem.Store, error) {
	s := &mem.Store{}
	if store == nil {
		return s, nil
	}
	for _, from := range store.Systems {
		to := &mem.System{X: from.X, Y: from.Y, Z: from.Z}
		switch from.Kind {
		case "Empty":
			to.Kind = mem.SKEmpty
		case "Blue Super Giant":
			to.Kind = mem.SKBlueSuperGiant
		case "Dense Dust Cloud":
			to.Kind = mem.SKDenseDustCloud
		case "Medium Dust Cloud":
			to.Kind = mem.SKMediumDustCloud
		case "Yellow Main Sequence":
			to.Kind = mem.SKYellowMainSequence
		default:
			to.Kind = mem.SKEmpty
		}
		s.Systems = append(s.Systems, to)
	}
	return s, nil
}

// StoreToJSDB converts an in-memory store to a JSDB store.
func StoreToJSDB(s *mem.Store) (*jsdb.Store, error) {
	store := &jsdb.Store{}
	store.Meta.Version = 1
	if s == nil {
		return store, nil
	}
	for _, from := range s.Systems {
		to := &jsdb.System{X: from.X, Y: from.Y, Z: from.Z}
		switch from.Kind {
		case mem.SKEmpty:
			to.Kind = "Empty"
		case mem.SKBlueSuperGiant:
			to.Kind = "Blue Super Giant"
		case mem.SKDenseDustCloud:
			to.Kind = "Dense Dust Cloud"
		case mem.SKMediumDustCloud:
			to.Kind = "Medium Dust Cloud"
		case mem.SKYellowMainSequence:
			to.Kind = "Yellow Main Sequence"
		default:
			to.Kind = ""
		}
		store.Systems = append(store.Systems, to)
	}
	return store, nil
}
