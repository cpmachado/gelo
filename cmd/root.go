/* Copyright © 2024 Carlos Pinto Machado<cpmachado@protonmail.com> */
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gelo",
	Short: "Retrieves, parses and generates a csv with data from FIDE rating List",
	Long: `gelo simply retrieves the last xml list from FIDE and outputs a csv
version of it, which greatly reduces the size of the file and makes
parsing easier.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
