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

package mem

// ideally, adapters would be friends in their own space,
// but Go doesn't allow for that, so they're here.

import "github.com/mdhender/lutymaps/store/jsdb"

// AdaptJSDBToStore converts a JSDB store to an in-memory store.
func AdaptJSDBToStore(store *jsdb.Store) (*Store, error) {
	s := &Store{}
	if store == nil {
		return s, nil
	}
	for _, from := range store.Systems {
		to := &System{x: from.X, y: from.Y, z: from.Z}
		switch from.Kind {
		case "Empty":
			to.kind = SKEmpty
		case "Blue Super Giant":
			to.kind = SKBlueSuperGiant
		case "Dense Dust Cloud":
			to.kind = SKDenseDustCloud
		case "Medium Dust Cloud":
			to.kind = SKMediumDustCloud
		case "Yellow Main Sequence":
			to.kind = SKYellowMainSequence
		default:
			to.kind = SKEmpty
		}
		s.systems = append(s.systems, to)
	}
	return s, nil
}

// AdaptStoreToJSDB converts an in-memory store to a JSDB store.
func AdaptStoreToJSDB(s *Store) (*jsdb.Store, error) {
	store := &jsdb.Store{}
	store.Meta.Version = 1
	if s == nil {
		return store, nil
	}
	for _, from := range s.systems {
		to := &jsdb.System{X: from.x, Y: from.y, Z: from.z}
		switch from.kind {
		case SKEmpty:
			to.Kind = "Empty"
		case SKBlueSuperGiant:
			to.Kind = "Blue Super Giant"
		case SKDenseDustCloud:
			to.Kind = "Dense Dust Cloud"
		case SKMediumDustCloud:
			to.Kind = "Medium Dust Cloud"
		case SKYellowMainSequence:
			to.Kind = "Yellow Main Sequence"
		default:
			to.Kind = ""
		}
		store.Systems = append(store.Systems, to)
	}
	return store, nil
}
