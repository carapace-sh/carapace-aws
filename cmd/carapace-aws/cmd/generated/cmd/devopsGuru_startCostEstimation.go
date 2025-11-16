package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_startCostEstimationCmd = &cobra.Command{
	Use:   "start-cost-estimation",
	Short: "Starts the creation of an estimate of the monthly cost to analyze your Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_startCostEstimationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_startCostEstimationCmd).Standalone()

		devopsGuru_startCostEstimationCmd.Flags().String("client-token", "", "The idempotency token used to identify each cost estimate request.")
		devopsGuru_startCostEstimationCmd.Flags().String("resource-collection", "", "The collection of Amazon Web Services resources used to create a monthly DevOps Guru cost estimate.")
		devopsGuru_startCostEstimationCmd.MarkFlagRequired("resource-collection")
	})
	devopsGuruCmd.AddCommand(devopsGuru_startCostEstimationCmd)
}
