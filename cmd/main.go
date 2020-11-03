package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "app",
		Short: "Kumparan commands",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "run-app",
		Short: "Run Kumparan endpoints",
		Run: func(cmd *cobra.Command, args []string) {
			envInit()
			app()
		}})

	cmd.AddCommand(&cobra.Command{
		Use:   "run-nsq",
		Short: "Run Kumparan NSQ Consumer",
		Run: func(cmd *cobra.Command, args []string) {
			envInit()
			nsqCmd()
		}})

	cmd.Execute()
}

func envInit() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
