package cmd

import (
	"log"

	"github.com/radavis/et/horizon"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatal("missing slug name")
		}

		slug := args[0]
		lesson := horizon.GetLesson(slug)
		// fmt.Printf("%v\n", lesson.ArchiveUrl)
		lesson.GetArchive()
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
