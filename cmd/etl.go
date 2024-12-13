package cmd

import (
	gelo "github.com/cpmachado/gelo/lib"
	"github.com/spf13/cobra"
)

// etlCmd represents the etl command
var etlCmd = &cobra.Command{
	Use:   "etl",
	Short: "Extract, parse XML from fide to csv",
	Long:  `Extract, parse XML from fide to csv`,
	Run: func(cmd *cobra.Command, args []string) {
		gelo.ExtractAndGenerateCsv()
	},
}

func init() {
	rootCmd.AddCommand(etlCmd)
}
