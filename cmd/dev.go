package cmd

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"

	"go-serverless-demo/internal/api"
	"go-serverless-demo/internal/db"
	"go-serverless-demo/internal/echo"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run server locally",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tableName := "go-todo"
		d := db.New(&aws.Config{
			Endpoint: aws.String("http://localhost:8000"),
			Region:   aws.String("ap-northeast-2"),
		})
		err := d.CreateTable(tableName, new(db.Todo))
		if err != nil {
			log.Println(err)
		}

		d.SetTable(tableName)

		a := api.NewAPI(d)
		e := echo.NewEcho(a)

		e.Start(":8080")
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
