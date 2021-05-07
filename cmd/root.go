/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var host string
var port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pcheck",
	Short:   "\"pcheck\" check certain TCP port is opend",
	Example: "pcheck -i 192.168.0.10 -p 8888",
	Run: func(cmd *cobra.Command, args []string) {
		address := net.JoinHostPort(host, port)
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			fmt.Println(err)
		} else if conn != nil {
			defer conn.Close()
			fmt.Printf("%s:%s is opened \n", host, port)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&host, "host", "i", "", "destination address for checking (required)")
	rootCmd.MarkFlagRequired("host")
	rootCmd.Flags().StringVarP(&port, "port", "p", "", "destination port for checking (required)")
	rootCmd.MarkFlagRequired("port")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.port-checker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".port-checker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".port-checker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
