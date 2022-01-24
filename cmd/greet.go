/*
Copyright © 2022 Gavin Campbell <gavin@gavincampbell.dev>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rsc.io/quote"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet the user",
	Long: `This is a code example for a post on gavincampbell.dev`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(quote.Hello())
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
