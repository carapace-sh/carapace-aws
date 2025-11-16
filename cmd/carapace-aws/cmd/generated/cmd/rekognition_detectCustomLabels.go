package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_detectCustomLabelsCmd = &cobra.Command{
	Use:   "detect-custom-labels",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_detectCustomLabelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_detectCustomLabelsCmd).Standalone()

		rekognition_detectCustomLabelsCmd.Flags().String("image", "", "")
		rekognition_detectCustomLabelsCmd.Flags().String("max-results", "", "Maximum number of results you want the service to return in the response.")
		rekognition_detectCustomLabelsCmd.Flags().String("min-confidence", "", "Specifies the minimum confidence level for the labels to return.")
		rekognition_detectCustomLabelsCmd.Flags().String("project-version-arn", "", "The ARN of the model version that you want to use.")
		rekognition_detectCustomLabelsCmd.MarkFlagRequired("image")
		rekognition_detectCustomLabelsCmd.MarkFlagRequired("project-version-arn")
	})
	rekognitionCmd.AddCommand(rekognition_detectCustomLabelsCmd)
}
