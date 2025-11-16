package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeClusterSchedulerConfigCmd = &cobra.Command{
	Use:   "describe-cluster-scheduler-config",
	Short: "Description of the cluster policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeClusterSchedulerConfigCmd).Standalone()

	sagemaker_describeClusterSchedulerConfigCmd.Flags().String("cluster-scheduler-config-id", "", "ID of the cluster policy.")
	sagemaker_describeClusterSchedulerConfigCmd.Flags().String("cluster-scheduler-config-version", "", "Version of the cluster policy.")
	sagemaker_describeClusterSchedulerConfigCmd.MarkFlagRequired("cluster-scheduler-config-id")
	sagemakerCmd.AddCommand(sagemaker_describeClusterSchedulerConfigCmd)
}
