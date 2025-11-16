package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_getCostEstimationCmd = &cobra.Command{
	Use:   "get-cost-estimation",
	Short: "Returns an estimate of the monthly cost for DevOps Guru to analyze your Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_getCostEstimationCmd).Standalone()

	devopsGuru_getCostEstimationCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuruCmd.AddCommand(devopsGuru_getCostEstimationCmd)
}
