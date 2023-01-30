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

package cli

import (
	"github.com/mdhender/lutymaps/pkg/adapters"
	"github.com/mdhender/lutymaps/pkg/server"
	"github.com/mdhender/lutymaps/pkg/stores/jsdb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Serve data for the engine",
	Long:  `Provide a REST-ish API for engine data.`,
	Run: func(cmd *cobra.Command, args []string) {
		jsPath := "galaxy-001.json"
		jstore, err := jsdb.New(jsPath)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("serve: loaded %q\n", jsPath)
		jsPath = "accounts.json"
		jsAccts := jsdb.AccountStore{}
		if err = jsAccts.Load(jsPath); err != nil {
			log.Fatal(err)
		}
		log.Printf("serve: loaded %q\n", jsPath)

		mstore, err := adapters.JSDBToStore(jstore)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("serve: adapted galaxy store\n")
		mstore.Accounts, err = adapters.JSAccountsToMemAccounts(jsAccts)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("serve: adapted accounts store\n")

		s, err := server.New(server.WithAuthentication(mstore), server.WithAuthorization(mstore))
		if err != nil {
			log.Fatal(err)
		}
		router := s.Routes()

		log.Printf("server: listening on %q\n", net.JoinHostPort(cliConfig.Server.Host, cliConfig.Server.Port))
		log.Fatal(http.ListenAndServe(net.JoinHostPort(cliConfig.Server.Host, cliConfig.Server.Port), router))
	},
}

func init() {
	cmdMain.AddCommand(cmdServe)
	cmdServe.Flags().StringVar(&cliConfig.Server.Host, "host", "", "interface to run server on")
	_ = viper.BindPFlag("host", cmdServe.Flags().Lookup("host"))
	cmdServe.Flags().StringVarP(&cliConfig.Server.Port, "port", "p", "3000", "port to run server on")
	_ = viper.BindPFlag("port", cmdServe.Flags().Lookup("port"))
}
