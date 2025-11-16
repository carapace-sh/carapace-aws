package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "You can use the Billing API to programatically list the billing views available to you for a given time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingCmd).Standalone()

	rootCmd.AddCommand(billingCmd)
}
