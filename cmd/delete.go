/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var id string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a expense",
	Long: `Delete an expense by providing the expense ID.
Example:
expense-tracker delete --delete YOUR_EXPENSE_ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			fmt.Println("Enter a id to delete!!!!")
			return
		}
		index := -1
		for i, v := range Expenses {
			if id == v.Id {
				index = i
			}
		}
		if index == -1 {
			fmt.Println("NO id present")
			return
		}
		Expenses = append(Expenses[:index], Expenses[index+1:]...)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringVarP(&id, "delete", "d", "", "Delete a expense")
	deleteCmd.MarkFlagRequired("delete")
}
