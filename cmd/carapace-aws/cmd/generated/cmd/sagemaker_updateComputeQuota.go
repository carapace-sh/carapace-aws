package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateComputeQuotaCmd = &cobra.Command{
	Use:   "update-compute-quota",
	Short: "Update the compute allocation definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateComputeQuotaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateComputeQuotaCmd).Standalone()

		sagemaker_updateComputeQuotaCmd.Flags().String("activation-state", "", "The state of the compute allocation being described.")
		sagemaker_updateComputeQuotaCmd.Flags().String("compute-quota-config", "", "Configuration of the compute allocation definition.")
		sagemaker_updateComputeQuotaCmd.Flags().String("compute-quota-id", "", "ID of the compute allocation definition.")
		sagemaker_updateComputeQuotaCmd.Flags().String("compute-quota-target", "", "The target entity to allocate compute resources to.")
		sagemaker_updateComputeQuotaCmd.Flags().String("description", "", "Description of the compute allocation definition.")
		sagemaker_updateComputeQuotaCmd.Flags().String("target-version", "", "Target version.")
		sagemaker_updateComputeQuotaCmd.MarkFlagRequired("compute-quota-id")
		sagemaker_updateComputeQuotaCmd.MarkFlagRequired("target-version")
	})
	sagemakerCmd.AddCommand(sagemaker_updateComputeQuotaCmd)
}
