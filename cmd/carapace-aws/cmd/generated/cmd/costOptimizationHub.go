package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHubCmd = &cobra.Command{
	Use:   "cost-optimization-hub",
	Short: "You can use the Cost Optimization Hub API to programmatically identify, filter, aggregate, and quantify savings for your cost optimization recommendations across multiple Amazon Web Services Regions and Amazon Web Services accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHubCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(costOptimizationHubCmd).Standalone()

	})
	rootCmd.AddCommand(costOptimizationHubCmd)
}
