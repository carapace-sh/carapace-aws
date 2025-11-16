package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgetsCmd = &cobra.Command{
	Use:   "budgets",
	Short: "Use the Amazon Web Services Budgets API to plan your service usage, service costs, and instance reservations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgetsCmd).Standalone()

	})
	rootCmd.AddCommand(budgetsCmd)
}
