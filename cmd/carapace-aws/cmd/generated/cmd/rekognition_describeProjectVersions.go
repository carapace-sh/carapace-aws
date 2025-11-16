package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_describeProjectVersionsCmd = &cobra.Command{
	Use:   "describe-project-versions",
	Short: "Lists and describes the versions of an Amazon Rekognition project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_describeProjectVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_describeProjectVersionsCmd).Standalone()

		rekognition_describeProjectVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
		rekognition_describeProjectVersionsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more results to retrieve), Amazon Rekognition returns a pagination token in the response.")
		rekognition_describeProjectVersionsCmd.Flags().String("project-arn", "", "The Amazon Resource Name (ARN) of the project that contains the model/adapter you want to describe.")
		rekognition_describeProjectVersionsCmd.Flags().String("version-names", "", "A list of model or project version names that you want to describe.")
		rekognition_describeProjectVersionsCmd.MarkFlagRequired("project-arn")
	})
	rekognitionCmd.AddCommand(rekognition_describeProjectVersionsCmd)
}
