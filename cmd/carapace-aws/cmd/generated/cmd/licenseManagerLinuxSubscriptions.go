package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptionsCmd = &cobra.Command{
	Use:   "license-manager-linux-subscriptions",
	Short: "With License Manager, you can discover and track your commercial Linux subscriptions on running Amazon EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerLinuxSubscriptionsCmd).Standalone()

	})
	rootCmd.AddCommand(licenseManagerLinuxSubscriptionsCmd)
}
