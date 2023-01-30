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
	"github.com/mdhender/lutymaps/pkg/scan"
	"github.com/mdhender/lutymaps/pkg/stores/jsdb"
	"github.com/mdhender/lutymaps/pkg/stores/mem"
	"github.com/spf13/cobra"
	"log"
)

var cmdScan = &cobra.Command{
	Use:   "scan",
	Short: "Scan a sector to a PNG file",
	Long:  `Create an image from a sector scan.`,
	Run: func(cmd *cobra.Command, args []string) {
		jstore, err := jsdb.New("galaxy-001.json")
		if err != nil {
			log.Fatal(err)
		}

		mstore, err := adapters.JSDBToStore(jstore)
		if err != nil {
			log.Fatal(err)
		}

		//jstore, err = adapters.StoreToJSDB(mstore)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//kinds := []string{
		//	"Blue Super Giant",
		//	"Blue Super Giant",
		//	"Blue Super Giant",
		//	"Yellow Main Sequence",
		//	"Yellow Main Sequence",
		//	"Dense Dust Cloud",
		//	"Medium Dust Cloud",
		//	"Light Dust Cloud",
		//}
		//for i := len(jstore.Systems); i < 10_000; i++ {
		//	jstore.Systems = append(jstore.Systems, &jsdb.System{
		//		X:    rand.Intn(51) - 25,
		//		Y:    rand.Intn(51) - 25,
		//		Z:    rand.Intn(51) - 25,
		//		Kind: kinds[rand.Intn(len(kinds))],
		//	})
		//}
		//err = jstore.Save("galaxy-002.json")
		//if err != nil {
		//	log.Fatal(err)
		//}

		err = scan.New(mstore, mem.FilterBySector(0, 0, 0, 50), "scan.png")
		if err != nil {
			log.Fatal(err)
		}

		//examples1.Beads()
		//examples1.Example1()
		//examples1.Outline()
		//examples1.SliceBowser()
		//examples1.SliceSuzanne()
		//examples1.Slices()
		//examples1.Suzanne()
		//examples1.VoxelizeBowser()
		//examples1.VoxelizeBunny()

		//earth.Main()
		//shapes.Main()
		//teapot.Main()
	},
}

func init() {
	cmdMain.AddCommand(cmdScan)
}
