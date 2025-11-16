package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteProjectVersionCmd = &cobra.Command{
	Use:   "delete-project-version",
	Short: "Deletes a Rekognition project model or project version, like a Amazon Rekognition Custom Labels model or a custom adapter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteProjectVersionCmd).Standalone()

	rekognition_deleteProjectVersionCmd.Flags().String("project-version-arn", "", "The Amazon Resource Name (ARN) of the project version that you want to delete.")
	rekognition_deleteProjectVersionCmd.MarkFlagRequired("project-version-arn")
	rekognitionCmd.AddCommand(rekognition_deleteProjectVersionCmd)
}
