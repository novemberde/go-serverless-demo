package cmd

import (
	"go-serverless-demo/internal/api"
	"log"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run server locally",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(api.New().Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
