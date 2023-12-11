/*
Copyright © 2023 Foolin
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5 file [file2] [file3] ...",
	Short: "Calculate the hexadecimal value of the MD5 files",
	Long:  `Calculate the hexadecimal value of the MD5 files`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// md5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
