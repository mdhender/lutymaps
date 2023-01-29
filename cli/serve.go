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
	"fmt"
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
		fmt.Printf("listening on %q using router\n", net.JoinHostPort(cliConfig.Server.Host, cliConfig.Server.Port))
		log.Fatal(http.ListenAndServe(net.JoinHostPort(cliConfig.Server.Host, cliConfig.Server.Port), nil))
	},
}

func init() {
	cmdMain.AddCommand(cmdServe)
	cmdServe.Flags().StringVar(&cliConfig.Server.Host, "host", "", "interface to run server on")
	_ = viper.BindPFlag("host", cmdServe.Flags().Lookup("host"))
	cmdServe.Flags().StringVarP(&cliConfig.Server.Port, "port", "p", "3000", "port to run server on")
	_ = viper.BindPFlag("port", cmdServe.Flags().Lookup("port"))
}
