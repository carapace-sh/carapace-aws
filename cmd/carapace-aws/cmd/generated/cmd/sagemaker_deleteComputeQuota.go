package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteComputeQuotaCmd = &cobra.Command{
	Use:   "delete-compute-quota",
	Short: "Deletes the compute allocation from the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteComputeQuotaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteComputeQuotaCmd).Standalone()

		sagemaker_deleteComputeQuotaCmd.Flags().String("compute-quota-id", "", "ID of the compute allocation definition.")
		sagemaker_deleteComputeQuotaCmd.MarkFlagRequired("compute-quota-id")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteComputeQuotaCmd)
}
