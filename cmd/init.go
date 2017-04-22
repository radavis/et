package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/mitchellh/go-homedir"
	. "github.com/radavis/et/common"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize current directory as a work area",
	Run: func(cmd *cobra.Command, args []string) {
		username := StringRequired("Username: ")
		token := StringRequired("Token: ")

		configTemplate := fmt.Sprintf(`username: %s
token: %s
`, username, token)
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		configPath := path.Join(home, ".et.yaml")

		err = ioutil.WriteFile(configPath, []byte(configTemplate), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Saved configuration to %s\n", configPath)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
