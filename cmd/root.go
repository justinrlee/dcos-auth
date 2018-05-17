// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
//	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
)

var privateKeyFile string
var uid string
var validTime int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dcos-auth",
	Short: "Helper tool to authenticate against DC/OS Enterprise ACS",
	Long:  `Helper tool to authenticate against DC/OS Enterprise ACS`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&privateKeyFile, "privatekeyfile", "k", "", "Private Key File")
	rootCmd.PersistentFlags().StringVarP(&uid, "uid", "u", "", "User Id")
	rootCmd.PersistentFlags().IntVarP(&validTime, "time", "t", 10, "Token validity (in seconds)")

}