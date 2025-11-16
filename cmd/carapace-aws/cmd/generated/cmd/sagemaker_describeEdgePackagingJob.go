package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeEdgePackagingJobCmd = &cobra.Command{
	Use:   "describe-edge-packaging-job",
	Short: "A description of edge packaging jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeEdgePackagingJobCmd).Standalone()

	sagemaker_describeEdgePackagingJobCmd.Flags().String("edge-packaging-job-name", "", "The name of the edge packaging job.")
	sagemaker_describeEdgePackagingJobCmd.MarkFlagRequired("edge-packaging-job-name")
	sagemakerCmd.AddCommand(sagemaker_describeEdgePackagingJobCmd)
}
