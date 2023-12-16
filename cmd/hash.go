/*
Copyright © 2023 Foolin
*/
package cmd

import (
	"github.com/foolin/sumdiff"
	"github.com/foolin/sumdiff/internal/plog"
	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/foolin/sumdiff/internal/vo"
	"github.com/spf13/cobra"
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash <algorithm> <path> ...",
	Short: "Calculate hash algorithm [md5|sha1|sha256|sha512] hex string",
	Long:  `Calculate hash algorithm [md5|sha1|sha256|sha512] hex string, eg: hash sha1 a.text`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		results, err := sumdiff.HashWithArgs(args...)
		statusbar.Clean()
		if err != nil {
			plog.Writeln(err)
			return
		}
		plog.WriteTable(vo.HashToTable(results))
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
