/*
Copyright © 2023 Foolin
*/
package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/foolin/sumdiff/internal/app"
	"github.com/foolin/sumdiff/internal/plog"
	"github.com/foolin/sumdiff/vo"
	"github.com/spf13/cobra"
	"hash"
	"strings"
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash [md5|sha1|sha256|sha512] path path2 ...",
	Short: "Calculate file hash hex value",
	Long:  `Calculate file hash hex value`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		t := strings.ToLower(args[0])
		var hash hash.Hash
		switch t {
		case "md5":
			hash = md5.New()
		case "sha1":
			hash = sha1.New()
		case "sha256":
			hash = sha256.New()
		case "sha512":
			hash = sha512.New()
		default:
			plog.PrintEnd("not support hash=%v", t)
			return
		}
		results, err := app.Hash(hash, args[1:]...)
		if err != nil {
			plog.PrintEnd("hash error: %v", err)
			return
		}
		plog.PrintTable(vo.HashToTable(results))
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
