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
	"crypto/sha1"
	"github.com/foolin/sumdiff/vo"

	"github.com/foolin/sumdiff"
	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/spf13/cobra"
)

// sha1Cmd represents the sha1 command
var sha1Cmd = &cobra.Command{
	Use:   "sha1 <path> ...",
	Short: "Calculate sha1 hex string",
	Long:  `Calculate sha1 hex string, eg: hash sha1 a.text`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		statusbar.Start()
		list, err := sumdiff.Hash(sha1.New(), args...)
		statusbar.Stop()
		if err != nil {
			writer.MustWrite(vo.NewErrInfo(err))
			return
		}
		writer.MustWrite(list)
	},
}

func init() {
	rootCmd.AddCommand(sha1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
