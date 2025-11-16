package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeBatchSegmentJobCmd = &cobra.Command{
	Use:   "describe-batch-segment-job",
	Short: "Gets the properties of a batch segment job including name, Amazon Resource Name (ARN), status, input and output configurations, and the ARN of the solution version used to generate segments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeBatchSegmentJobCmd).Standalone()

	personalize_describeBatchSegmentJobCmd.Flags().String("batch-segment-job-arn", "", "The ARN of the batch segment job to describe.")
	personalize_describeBatchSegmentJobCmd.MarkFlagRequired("batch-segment-job-arn")
	personalizeCmd.AddCommand(personalize_describeBatchSegmentJobCmd)
}
