/*
Copyright © 2020 Amanda Zhu <amandazhuyilan@gmail.com>

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
	"os/user"
	"runtime"

	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Displays current user's name and username",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "windows" {
			fmt.Println("Can't Execute this on a windows machine")
		} else {
			user := getUser()
			username := getUserName(user)
			name := getName(user)

			fmt.Println("Hello " + name + "!")
			fmt.Println("Username: " + username)
		}
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getUser() *user.User {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user
}

func getUserName(user *user.User) string {
	return user.Username
}

func getName(user *user.User) string {
	return user.Name
}
