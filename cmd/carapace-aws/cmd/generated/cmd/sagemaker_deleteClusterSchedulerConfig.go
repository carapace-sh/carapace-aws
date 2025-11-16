package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteClusterSchedulerConfigCmd = &cobra.Command{
	Use:   "delete-cluster-scheduler-config",
	Short: "Deletes the cluster policy of the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteClusterSchedulerConfigCmd).Standalone()

	sagemaker_deleteClusterSchedulerConfigCmd.Flags().String("cluster-scheduler-config-id", "", "ID of the cluster policy.")
	sagemaker_deleteClusterSchedulerConfigCmd.MarkFlagRequired("cluster-scheduler-config-id")
	sagemakerCmd.AddCommand(sagemaker_deleteClusterSchedulerConfigCmd)
}
