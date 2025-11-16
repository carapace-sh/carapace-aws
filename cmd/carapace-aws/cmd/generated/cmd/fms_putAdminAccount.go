package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_putAdminAccountCmd = &cobra.Command{
	Use:   "put-admin-account",
	Short: "Creates or updates an Firewall Manager administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_putAdminAccountCmd).Standalone()

	fms_putAdminAccountCmd.Flags().String("admin-account", "", "The Amazon Web Services account ID to add as an Firewall Manager administrator account.")
	fms_putAdminAccountCmd.Flags().String("admin-scope", "", "Configures the resources that the specified Firewall Manager administrator can manage.")
	fms_putAdminAccountCmd.MarkFlagRequired("admin-account")
	fmsCmd.AddCommand(fms_putAdminAccountCmd)
}
