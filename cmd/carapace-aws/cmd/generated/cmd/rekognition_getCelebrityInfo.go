package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getCelebrityInfoCmd = &cobra.Command{
	Use:   "get-celebrity-info",
	Short: "Gets the name and additional information about a celebrity based on their Amazon Rekognition ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getCelebrityInfoCmd).Standalone()

	rekognition_getCelebrityInfoCmd.Flags().String("id", "", "The ID for the celebrity.")
	rekognition_getCelebrityInfoCmd.MarkFlagRequired("id")
	rekognitionCmd.AddCommand(rekognition_getCelebrityInfoCmd)
}
