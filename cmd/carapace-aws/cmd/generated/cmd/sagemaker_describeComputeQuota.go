package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeComputeQuotaCmd = &cobra.Command{
	Use:   "describe-compute-quota",
	Short: "Description of the compute allocation definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeComputeQuotaCmd).Standalone()

	sagemaker_describeComputeQuotaCmd.Flags().String("compute-quota-id", "", "ID of the compute allocation definition.")
	sagemaker_describeComputeQuotaCmd.Flags().String("compute-quota-version", "", "Version of the compute allocation definition.")
	sagemaker_describeComputeQuotaCmd.MarkFlagRequired("compute-quota-id")
	sagemakerCmd.AddCommand(sagemaker_describeComputeQuotaCmd)
}
