package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listMobileDeviceAccessRulesCmd = &cobra.Command{
	Use:   "list-mobile-device-access-rules",
	Short: "Lists the mobile device access rules for the specified WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listMobileDeviceAccessRulesCmd).Standalone()

	workmail_listMobileDeviceAccessRulesCmd.Flags().String("organization-id", "", "The WorkMail organization for which to list the rules.")
	workmail_listMobileDeviceAccessRulesCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listMobileDeviceAccessRulesCmd)
}
