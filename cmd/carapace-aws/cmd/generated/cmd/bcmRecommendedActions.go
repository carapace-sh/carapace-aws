package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmRecommendedActionsCmd = &cobra.Command{
	Use:   "bcm-recommended-actions",
	Short: "You can use the Billing and Cost Management Recommended Actions API to programmatically query your best practices and recommendations to optimize your costs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmRecommendedActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmRecommendedActionsCmd).Standalone()

	})
	rootCmd.AddCommand(bcmRecommendedActionsCmd)
}
