package cmd

import (
	"fmt"

	"github.com/radavis/et/horizon"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		lessons := horizon.GetLessons()
		for i, l := range lessons.Lessons {
			fmt.Printf("%d  %s - %s\n", i, l.Title, l.Slug)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
