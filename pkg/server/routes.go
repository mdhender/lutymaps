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

package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"path/filepath"
	"strings"
)

func (s *Server) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET", "PUT", "POST", "DELETE", "HEAD", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept", "Accept-Encoding", "Accept-Language", "Authorization",
			"Cache-Control", "Connection", "Content-Type", "DNT", "Host",
			"Origin", "Pragma", "Referer", "User-Agent",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// public routes
	r.Mount("/", s.app.Router())

	// protected routes
	r.Route("/api", func(r chi.Router) {
		//r.Use(jwtauth.Verifier(tokenAuth)) // extract, verify, validate JWT
		////r.Use(jwtauth.Authenticator)       // handle valid and invalid JWT
		//r.Use(JWTAuthenticator)    // handle valid and invalid JWT
		r.Mount("/", s.api.Router()) // mount the api sub-router
	})

	// static files
	s.static = http.FileServer(http.Dir(s.public))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// osPath is the cleaned up path with our file system path separators
		osPath := filepath.Clean(r.RequestURI)
		// urlPath is the same but with url path separators
		urlPath := strings.ReplaceAll(osPath, "\\", "/")
		// verify that the request path in the url matches the cleaned up path
		if r.RequestURI != urlPath { // forbid when there's a mismatch
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		s.static.ServeHTTP(w, r)
	})

	s.router = r
	return r
}

// staticFiles serves static files.
func (s *Server) staticFiles(router *chi.Mux) {
}
