/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/zach-monroe/doit/todo"

	"github.com/spf13/cobra"
)

var priority int

func addRun(cmd *cobra.Command, args []string) {
	items := []todo.Item{}

	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}
	err := todo.SaveItems(dataFile, items)
	if err != nil {
		fmt.Printf("%v", err)
	}

}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "this allows you to add to new items to your todo list!",
	Long:  `add your todo item by running doit add <item>`,
	Run:   addRun,
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
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
