/*
Copyright © 2023 Foolin
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sha1Cmd represents the sha1 command
var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sha1 called")
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
