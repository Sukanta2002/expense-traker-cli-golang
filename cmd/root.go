/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var PATH = "./expense.json"
var Expenses []Expense

type Expense struct {
	Id          string  `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-traker-cli",
	Short: "A Simple expense tracker in your terminal",
	Long: `Expense Tracker CLI is a tool to manage your personal finances right from the terminal.
It allows you to add, view, and categorize expenses with ease.

Features include:
- Adding new expenses with a description and amount.
- Categorizing expenses for better tracking.
- Viewing a summary of expenses over different periods.

Get started quickly by using the available commands and flags.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		createFile()
		getData()
		defer SaveFile()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		createFile()
		getData()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		SaveFile()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.expense-traker-cli-golang.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createFile() {
	file, e := os.Open(PATH)
	if e != nil {
		os.Create(PATH)
	}
	file.Close()
}

func getData() {
	data, _ := os.ReadFile(PATH)
	if len(data) == 0 {
		return
	}
	json.Unmarshal(data, &Expenses)
}

func SaveFile() {
	data, e := json.MarshalIndent(Expenses, "", "\t")
	if e != nil {
		panic(e)
	}
	os.WriteFile(PATH, data, os.ModeAppend)
}
