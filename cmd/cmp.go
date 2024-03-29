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
	"github.com/foolin/sumdiff"
	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/foolin/sumdiff/internal/vlog"
	"github.com/spf13/cobra"
)

// cmpCmd represents the cmp command
var cmpCmd = &cobra.Command{
	Use:   "cmp",
	Short: "Compare two files or directories for a list of differences",
	Long:  `Compare two files or directories for a list of differences`,
	Run: func(cmd *cobra.Command, args []string) {
		statusbar.Start()
		ok, list, err := sumdiff.Cmp(args[0], args[1])
		statusbar.Stop()
		if err != nil {
			vlog.Printf("Error: %v", err)
		}
		vlog.Printf("Result: %v", ok)
		writer.MustWrite(list)
	},
}

func init() {
	rootCmd.AddCommand(cmpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//cmpCmd.PersistentFlags().BoolVarP(&detail, "detail", "", false, "Show detail result")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
