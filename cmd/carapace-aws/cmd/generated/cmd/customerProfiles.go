package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfilesCmd = &cobra.Command{
	Use:   "customer-profiles",
	Short: "Amazon Connect Customer Profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfilesCmd).Standalone()

	})
	rootCmd.AddCommand(customerProfilesCmd)
}
