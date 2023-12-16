/*
Copyright © 2023 Foolin
*/
package cmd

import (
	"github.com/foolin/sumdiff"
	"github.com/foolin/sumdiff/internal/plog"
	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff <path1> <path2>",
	Short: "Compare whether two files or directory are different",
	Long:  `Compare whether two files or directory are different`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := sumdiff.Diff(args[0], args[1])
		statusbar.Clean()
		if err != nil {
			plog.Writeln(err.Error())
		}
		plog.Writeln(ret)
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
