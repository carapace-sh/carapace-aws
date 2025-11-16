package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeClusterEventCmd = &cobra.Command{
	Use:   "describe-cluster-event",
	Short: "Retrieves detailed information about a specific event for a given HyperPod cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeClusterEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeClusterEventCmd).Standalone()

		sagemaker_describeClusterEventCmd.Flags().String("cluster-name", "", "The name or Amazon Resource Name (ARN) of the HyperPod cluster associated with the event.")
		sagemaker_describeClusterEventCmd.Flags().String("event-id", "", "The unique identifier (UUID) of the event to describe.")
		sagemaker_describeClusterEventCmd.MarkFlagRequired("cluster-name")
		sagemaker_describeClusterEventCmd.MarkFlagRequired("event-id")
	})
	sagemakerCmd.AddCommand(sagemaker_describeClusterEventCmd)
}
