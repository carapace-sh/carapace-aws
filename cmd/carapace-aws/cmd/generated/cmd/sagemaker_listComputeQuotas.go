package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listComputeQuotasCmd = &cobra.Command{
	Use:   "list-compute-quotas",
	Short: "List the resource allocation definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listComputeQuotasCmd).Standalone()

	sagemaker_listComputeQuotasCmd.Flags().String("cluster-arn", "", "Filter for ARN of the cluster.")
	sagemaker_listComputeQuotasCmd.Flags().String("created-after", "", "Filter for after this creation time.")
	sagemaker_listComputeQuotasCmd.Flags().String("created-before", "", "Filter for before this creation time.")
	sagemaker_listComputeQuotasCmd.Flags().String("max-results", "", "The maximum number of compute allocation definitions to list.")
	sagemaker_listComputeQuotasCmd.Flags().String("name-contains", "", "Filter for name containing this string.")
	sagemaker_listComputeQuotasCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
	sagemaker_listComputeQuotasCmd.Flags().String("sort-by", "", "Filter for sorting the list by a given value.")
	sagemaker_listComputeQuotasCmd.Flags().String("sort-order", "", "The order of the list.")
	sagemaker_listComputeQuotasCmd.Flags().String("status", "", "Filter for status.")
	sagemakerCmd.AddCommand(sagemaker_listComputeQuotasCmd)
}
