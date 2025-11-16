package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startProjectVersionCmd = &cobra.Command{
	Use:   "start-project-version",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startProjectVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_startProjectVersionCmd).Standalone()

		rekognition_startProjectVersionCmd.Flags().String("max-inference-units", "", "The maximum number of inference units to use for auto-scaling the model.")
		rekognition_startProjectVersionCmd.Flags().String("min-inference-units", "", "The minimum number of inference units to use.")
		rekognition_startProjectVersionCmd.Flags().String("project-version-arn", "", "The Amazon Resource Name(ARN) of the model version that you want to start.")
		rekognition_startProjectVersionCmd.MarkFlagRequired("min-inference-units")
		rekognition_startProjectVersionCmd.MarkFlagRequired("project-version-arn")
	})
	rekognitionCmd.AddCommand(rekognition_startProjectVersionCmd)
}
