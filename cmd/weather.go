/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"log"
	"os"

	"github.com/briandowns/openweathermap"
	"github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Displays the given city's current weather condition",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		OpenWeatherAPIKey := getContentfromTextFile("open_weather.txt")

		currentWeather, err := openweathermap.NewCurrent("C", "EN", OpenWeatherAPIKey) // fahrenheit (imperial) with Russian output
		if err != nil {
			log.Fatalln(err)
		}

		// TODO: Handle bad city names
		currentWeather.CurrentByName(city)
		fmt.Println(currentWeather.Weather)
	},
}

var city string

func init() {
	rootCmd.AddCommand(weatherCmd)
	weatherCmd.Flags().StringVarP(&city, "city", "c", "", "City name (required)")
	weatherCmd.MarkFlagRequired("city")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getContentfromTextFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	dataStr := string(data)
	return dataStr
}

// func getCityName() string {

// }
