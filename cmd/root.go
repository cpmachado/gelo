package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.cpmachado.pt/gelo/internal/config"
)

var (
	runVersionFlag bool
	cfgFile        string = config.DefaultConfigFile
)

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
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&runVersionFlag, "version", "v", false, "Display version")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", config.DefaultConfigFile, "config file")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(config.DefaultConfigHome)
		viper.SetConfigType("json")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
