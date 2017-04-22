package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mholt/archiver"
	"github.com/radavis/et/horizon"
	"github.com/spf13/cobra"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		filename := "archive.tar.gz"
		dir, _ := os.Getwd()

		tmp := strings.Split(dir, "/")
		slug := tmp[len(tmp)-1]

		archiver.TarGz.Make(filename, []string{dir})
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		result := horizon.PostLesson(slug, file)
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(submitCmd)
}
