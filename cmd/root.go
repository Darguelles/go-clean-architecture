package cmd

import (
	"fmt"
	"go-clean-architecture/server"
	"os"

	"github.com/spf13/cobra"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	appConfigs server.Configs
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "go-clean-architecture",
		Short: "Starts the webserver",
		Long: `Starts the webserver. You can pass configuration through commandline or cli flags. For example:

go-clean-architecture --port 8000`,
		Run: func(cmd *cobra.Command, args []string) {
			appConfigs.Port = viper.GetString("port")
			appConfigs.LogLVL = viper.GetString("log")
			if err := server.Start(appConfigs); err != nil {
				panic(err)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-clean-architecture.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Application Configs
	rootCmd.Flags().StringP("port", "p", "8000", "server port (default 8000)")
	rootCmd.Flags().String("log", "INFO", "set logger level (default Info)")
	rootCmd.Flags().Bool("devMode", false, "run in development mode")

	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		panic(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-clean-architecture" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-clean-architecture")
	}
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("GCA")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
