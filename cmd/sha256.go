/*
Copyright © 2023 Foolin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"crypto/sha256"
	"github.com/foolin/sumdiff/internal/write"
	"github.com/foolin/sumdiff/vo"

	"github.com/foolin/sumdiff"
	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/spf13/cobra"
)

// sha256Cmd represents the sha256 command
var sha256Cmd = &cobra.Command{
	Use:   "sha256 <path> ...",
	Short: "Calculate sha256 hex string",
	Long:  `Calculate sha256 hex string, eg: hash sha1 a.text`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		statusbar.Start()
		list, err := sumdiff.Hash(sha256.New(), args...)
		statusbar.Stop()
		w := write.NewStd()
		if err != nil {
			w.MustWrite(vo.NewErrInfo(err))
			return
		}
		w.MustWrite(list)
	},
}

func init() {
	rootCmd.AddCommand(sha256Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha256Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha256Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
