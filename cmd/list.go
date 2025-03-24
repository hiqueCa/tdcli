/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hiqueCa/go_todo_cli/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing to-dos",
	Long: `List existing to-dos in your to-dos file. Usage example:
	
	tdcli list
	`,
	Run: list,
}

func list(cmd *cobra.Command, args []string) {
	todos, err := todo.List(TodosFileName)

	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, item := range todos {
		fmt.Println(item)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
