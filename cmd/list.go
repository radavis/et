package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/radavis/et/horizon"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		lessons := horizon.Get()
		fmt.Println(lessons)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
