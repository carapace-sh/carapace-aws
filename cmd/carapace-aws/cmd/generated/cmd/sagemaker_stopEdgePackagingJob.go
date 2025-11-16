package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopEdgePackagingJobCmd = &cobra.Command{
	Use:   "stop-edge-packaging-job",
	Short: "Request to stop an edge packaging job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopEdgePackagingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_stopEdgePackagingJobCmd).Standalone()

		sagemaker_stopEdgePackagingJobCmd.Flags().String("edge-packaging-job-name", "", "The name of the edge packaging job.")
		sagemaker_stopEdgePackagingJobCmd.MarkFlagRequired("edge-packaging-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_stopEdgePackagingJobCmd)
}
