/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hiqueCa/go_todo_cli/todo"
	"github.com/spf13/cobra"
)

const TodosFileName string = "todos.json"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add to-dos",
	Long: `Add to-dos to your to-dos file. Usage example:
	
	tdcli add "lorem" ipsum "dolor amet"
	`,
	Run: addTodo,
}

func addTodo(cmd *cobra.Command, args []string) {
	items, err := todo.List("todos.json")

	if err != nil {
		fmt.Errorf("%v", err)
	}

	for _, label := range args {
		newTodo := todo.Item{Label: label}

		items = append(items, newTodo)
	}

	err = todo.Save(TodosFileName, items)

	if err != nil {
		fmt.Errorf("%v", err)
	}
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
}
