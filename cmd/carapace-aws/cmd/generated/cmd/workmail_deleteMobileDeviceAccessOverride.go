package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteMobileDeviceAccessOverrideCmd = &cobra.Command{
	Use:   "delete-mobile-device-access-override",
	Short: "Deletes the mobile device access override for the given WorkMail organization, user, and device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteMobileDeviceAccessOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_deleteMobileDeviceAccessOverrideCmd).Standalone()

		workmail_deleteMobileDeviceAccessOverrideCmd.Flags().String("device-id", "", "The mobile device for which you delete the override.")
		workmail_deleteMobileDeviceAccessOverrideCmd.Flags().String("organization-id", "", "The WorkMail organization for which the access override will be deleted.")
		workmail_deleteMobileDeviceAccessOverrideCmd.Flags().String("user-id", "", "The WorkMail user for which you want to delete the override.")
		workmail_deleteMobileDeviceAccessOverrideCmd.MarkFlagRequired("device-id")
		workmail_deleteMobileDeviceAccessOverrideCmd.MarkFlagRequired("organization-id")
		workmail_deleteMobileDeviceAccessOverrideCmd.MarkFlagRequired("user-id")
	})
	workmailCmd.AddCommand(workmail_deleteMobileDeviceAccessOverrideCmd)
}
