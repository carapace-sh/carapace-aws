package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptionsCmd = &cobra.Command{
	Use:   "license-manager-user-subscriptions",
	Short: "With License Manager, you can create user-based subscriptions to utilize licensed software with a per user subscription fee on Amazon EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptionsCmd).Standalone()

	rootCmd.AddCommand(licenseManagerUserSubscriptionsCmd)
}
