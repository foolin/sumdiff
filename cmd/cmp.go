/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/foolin/sumdiff/internal/app"
	"github.com/foolin/sumdiff/internal/plog"
	"github.com/foolin/sumdiff/internal/vo"

	"github.com/spf13/cobra"
)

// cmpCmd represents the cmp command
var cmpCmd = &cobra.Command{
	Use:   "cmp",
	Short: "Comparing different files or directory",
	Long:  `Comparing different files or directory`,
	Run: func(cmd *cobra.Command, args []string) {
		_, list, err := app.Cmp(args[0], args[1])
		if err != nil {
			plog.Writeln(err.Error())
		}
		plog.WriteTable(vo.CmpToTable(list))
	},
}

func init() {
	rootCmd.AddCommand(cmpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
