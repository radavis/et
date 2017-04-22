package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// "github.com/mholt/archiver"
// "github.com/radavis/et/horizon"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("%v", args)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
