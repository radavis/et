package cmd

import (
	"fmt"

	"github.com/radavis/et/horizon"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		lesson := horizon.GetLesson("golf-score")
		fmt.Printf("%v\n", lesson.ArchiveUrl)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
