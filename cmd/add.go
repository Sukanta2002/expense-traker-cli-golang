/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var description string
var amount float64

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a expense",
	Long: `Add a expense, example:
  expense-tracker add --description "Lunch" --amount 20`,
	Run: func(cmd *cobra.Command, args []string) {
		if amount == 0 && description == "" {
			fmt.Println("Please Enter a description and a amount!!!!")
			return
		}
		id := uuid.New().String()
		createDate := time.Now().Format(time.DateOnly)
		fmt.Println(createDate, id)
		data := Expense{Id: id, Date: createDate, Description: description, Amount: amount}
		Expenses = append(Expenses, data)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Expense description")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Expense amount")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
