package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_stopProjectVersionCmd = &cobra.Command{
	Use:   "stop-project-version",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_stopProjectVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_stopProjectVersionCmd).Standalone()

		rekognition_stopProjectVersionCmd.Flags().String("project-version-arn", "", "The Amazon Resource Name (ARN) of the model version that you want to stop.")
		rekognition_stopProjectVersionCmd.MarkFlagRequired("project-version-arn")
	})
	rekognitionCmd.AddCommand(rekognition_stopProjectVersionCmd)
}
