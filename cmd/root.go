package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var runVersionFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gelo",
	Short: "A tool to keep up with ratings",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !runVersionFlag {
			return cmd.Help()
		}
		progBuildInfo, ok := debug.ReadBuildInfo()
		if ok {
			fmt.Printf("%s-%s", cmd.Use, progBuildInfo.Main.Version)
		} else {
			fmt.Printf("%s-unknown", cmd.Use)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&runVersionFlag, "version", "v", false, "Display version")
}
