package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_createComputeEnvironmentCmd = &cobra.Command{
	Use:   "create-compute-environment",
	Short: "Creates an Batch compute environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_createComputeEnvironmentCmd).Standalone()

	batch_createComputeEnvironmentCmd.Flags().String("compute-environment-name", "", "The name for your compute environment.")
	batch_createComputeEnvironmentCmd.Flags().String("compute-resources", "", "Details about the compute resources managed by the compute environment.")
	batch_createComputeEnvironmentCmd.Flags().String("context", "", "Reserved.")
	batch_createComputeEnvironmentCmd.Flags().String("eks-configuration", "", "The details for the Amazon EKS cluster that supports the compute environment.")
	batch_createComputeEnvironmentCmd.Flags().String("service-role", "", "The full Amazon Resource Name (ARN) of the IAM role that allows Batch to make calls to other Amazon Web Services services on your behalf.")
	batch_createComputeEnvironmentCmd.Flags().String("state", "", "The state of the compute environment.")
	batch_createComputeEnvironmentCmd.Flags().String("tags", "", "The tags that you apply to the compute environment to help you categorize and organize your resources.")
	batch_createComputeEnvironmentCmd.Flags().String("type", "", "The type of the compute environment: `MANAGED` or `UNMANAGED`.")
	batch_createComputeEnvironmentCmd.Flags().String("unmanagedv-cpus", "", "The maximum number of vCPUs for an unmanaged compute environment.")
	batch_createComputeEnvironmentCmd.MarkFlagRequired("compute-environment-name")
	batch_createComputeEnvironmentCmd.MarkFlagRequired("type")
	batchCmd.AddCommand(batch_createComputeEnvironmentCmd)
}
