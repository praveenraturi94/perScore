// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"log"

	"perScore/app/model"
	"perScore/app/shared"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // ...
	"github.com/spf13/cobra"
)

// createDBCmd represents the createDB command
var createDBCmd = &cobra.Command{
	Use:   "createDB",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db1, err1 := gorm.Open("mysql", "root:root@/?charset=utf8&parseTime=True&loc=Local")
		defer db1.Close()

		if err1 != nil {
			log.Fatal(err1)
		}
		db1.Exec("CREATE DATABASE IF NOT EXISTS library DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci'")
		shared.DBCon, err = gorm.Open("mysql", "root:root@/library?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatal(err)
		}
		defer shared.DBCon.Close()
		model.SetupDatabase()
	},
}

func init() {
	RootCmd.AddCommand(createDBCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createDBCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createDBCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
