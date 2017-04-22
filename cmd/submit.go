package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/mholt/archiver"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		// fmt.Println("submit called")
		filename := "archive.tar.gz"
		dir, _ := os.Getwd()
		tmp := strings.Split("/", dir)
		slug := tmp[len(temp-1)]

		archiver.TarGz.Make(filename, []string{dir})
		os.Open(filename)
		horizon.PostLesson(slug, filename)
	},
}

func init() {
	RootCmd.AddCommand(submitCmd)
}
