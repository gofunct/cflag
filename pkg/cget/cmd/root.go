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
	"github.com/gofunct/cflag/pkg"
	"github.com/gofunct/gocfg"
	"os"

	"github.com/spf13/cobra"
)

var (
	config = gocfg.New(Conf.CfgFile, "cget")
)

type Config struct {
	CfgFile string
	Source string
	Dest string
}

var Conf = &Config{}

func init() {
	{
		rootCmd.AddCommand(fileCmd)
		rootCmd.AddCommand(dirCmd)
	}
	rootCmd.PersistentFlags().StringVarP(&Conf.CfgFile, "config", "c", os.Getenv("PWD")+".config.yaml", "config file (default is $PWD/.config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&Conf.Source, "source", "s", "", "url to download from")
	rootCmd.PersistentFlags().StringVarP(&Conf.Dest, "dest", "d", ".", "destination directory")
	config.BindCmd(rootCmd)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cget",
	Short: "download remote files or directories",
	Long: `Download from a number of different sources (file paths, Git, HTTP, Mercurial using only a URL as input.
If you want to download only a specific subdirectory from a downloaded directory, you can specify a 
subdirectory after a double-slash //
`,
}

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "download a directory from a remote url",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Getter("dir", Conf.Source, Conf.Dest)
	},
}

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "download a file from a remote url",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Getter("file", Conf.Source, Conf.Dest)
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
