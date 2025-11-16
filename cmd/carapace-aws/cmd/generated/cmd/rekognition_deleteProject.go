package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes a Amazon Rekognition project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_deleteProjectCmd).Standalone()

		rekognition_deleteProjectCmd.Flags().String("project-arn", "", "The Amazon Resource Name (ARN) of the project that you want to delete.")
		rekognition_deleteProjectCmd.MarkFlagRequired("project-arn")
	})
	rekognitionCmd.AddCommand(rekognition_deleteProjectCmd)
}
