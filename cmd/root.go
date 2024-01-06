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
	"os"
	"path/filepath"

	"github.com/foolin/sumdiff/internal/util"
	"github.com/foolin/sumdiff/internal/vlog"
	"github.com/foolin/sumdiff/internal/write"
	"github.com/spf13/cobra"
)

var config *Config
var writer *write.Writer
var file *os.File

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sumdiff",
	Short: "A useful comparison tool for differences",
	Long:  `A useful comparison tool for differences and hash`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//Verbose
		vlog.SetVerbose(config.Verbose)

		//Write
		format := write.Table
		if config.Format != "" {
			var ok bool
			format, ok = write.FormatOfName(config.Format)
			if !ok {
				vlog.Exit("Format invalid: %v\n", config.Format)
				return
			}
		}

		w := os.Stdout
		if config.Output != "" {
			path := util.FormatPath(config.Output)
			_ = os.MkdirAll(filepath.Dir(path), 0755)
			var err error
			file, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
			if err != nil {
				vlog.Exit("Open file %v error: %\n", path, err)
				return
			}
			w = file
		}

		//Create writer
		writer = write.New(w, format)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if file != nil {
			err := file.Close()
			if err != nil {
				vlog.Printf("Close file error: %v\n", err)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	config = &Config{
		Verbose: false,
		Format:  "table",
		Output:  "",
	}
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sumdiff.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "Verbose output info")
	rootCmd.PersistentFlags().StringVarP(&config.Format, "format", "f", "table", "Format: table|json|csv|yaml")
	rootCmd.PersistentFlags().StringVarP(&config.Output, "output", "o", "", "Output filename")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
