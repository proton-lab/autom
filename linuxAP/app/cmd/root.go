// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/proton-lab/autom/linuxAP/app/common"
	"github.com/proton-lab/autom/linuxAP/app/cmdservice"
	"log"
)

var remoteaddr string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "protonap",
	Short: "start protonap",
	Long: `start protonap`,
	//Uncomment the following line if your bare application
	//has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if b,err:=common.IsLinxAPProcessCanStarted();err!=nil{
			log.Println(err)
			return
		}else{
			if b{
				log.Println("protonap begin to start...")
			}else {
				return
			}
		}

		cmdinst:=cmdservice.GetCmdServerInst()
		cmdinst.StartCmdService()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&remoteaddr, "addr", "a","", "remote address")
}

//// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		// Search config in home directory with name ".app" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".app")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}