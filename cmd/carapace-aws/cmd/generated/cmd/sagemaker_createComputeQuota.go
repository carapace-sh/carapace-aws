package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createComputeQuotaCmd = &cobra.Command{
	Use:   "create-compute-quota",
	Short: "Create compute allocation definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createComputeQuotaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createComputeQuotaCmd).Standalone()

		sagemaker_createComputeQuotaCmd.Flags().String("activation-state", "", "The state of the compute allocation being described.")
		sagemaker_createComputeQuotaCmd.Flags().String("cluster-arn", "", "ARN of the cluster.")
		sagemaker_createComputeQuotaCmd.Flags().String("compute-quota-config", "", "Configuration of the compute allocation definition.")
		sagemaker_createComputeQuotaCmd.Flags().String("compute-quota-target", "", "The target entity to allocate compute resources to.")
		sagemaker_createComputeQuotaCmd.Flags().String("description", "", "Description of the compute allocation definition.")
		sagemaker_createComputeQuotaCmd.Flags().String("name", "", "Name to the compute allocation definition.")
		sagemaker_createComputeQuotaCmd.Flags().String("tags", "", "Tags of the compute allocation definition.")
		sagemaker_createComputeQuotaCmd.MarkFlagRequired("cluster-arn")
		sagemaker_createComputeQuotaCmd.MarkFlagRequired("compute-quota-config")
		sagemaker_createComputeQuotaCmd.MarkFlagRequired("compute-quota-target")
		sagemaker_createComputeQuotaCmd.MarkFlagRequired("name")
	})
	sagemakerCmd.AddCommand(sagemaker_createComputeQuotaCmd)
}
