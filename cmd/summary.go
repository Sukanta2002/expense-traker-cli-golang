/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var month int

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "View the Total expense of your",
	Long: `Example :
	expense-traker summary`,
	Run: func(cmd *cobra.Command, args []string) {
		var total float64 = 0
		if month == 0 {
			for _, v := range Expenses {
				total += v.Amount
			}
		} else if month > 12 {
			fmt.Println("Enter a number from 1 to 12")
		} else {
			for _, v := range Expenses {
				mon, _ := strconv.ParseInt(strings.Split(v.Date, "-")[1], 10, 64)
				if int64(month) == mon {
					total += v.Amount
				}
			}
		}

		fmt.Println("Total expenses: ", total)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "Summary of a month")
}
