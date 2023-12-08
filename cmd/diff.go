/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/foolin/sumdiff/internal/app"
	"github.com/foolin/sumdiff/internal/plog"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff path1 path2",
	Short: "Compare different files or directory",
	Long:  `Compare different files or directory`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := app.Diff(args[0], args[1])
		if err != nil {
			plog.Print(err.Error())
		} else {
			plog.Print(ret)
		}

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
