package cmd

import (
	"fmt"

	. "github.com/radavis/et/common"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize current directory as a work area",
	Run: func(cmd *cobra.Command, args []string) {
		username := StringRequired("Enter your username: ")
		token := StringRequired("Enter your ET token: ")
		host := StringRequired("Host for ET: ")
		fmt.Println(username)
		fmt.Println(token)
		fmt.Println(host)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
