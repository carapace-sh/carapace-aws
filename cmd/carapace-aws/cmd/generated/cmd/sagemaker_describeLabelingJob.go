package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeLabelingJobCmd = &cobra.Command{
	Use:   "describe-labeling-job",
	Short: "Gets information about a labeling job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeLabelingJobCmd).Standalone()

	sagemaker_describeLabelingJobCmd.Flags().String("labeling-job-name", "", "The name of the labeling job to return information for.")
	sagemaker_describeLabelingJobCmd.MarkFlagRequired("labeling-job-name")
	sagemakerCmd.AddCommand(sagemaker_describeLabelingJobCmd)
}
