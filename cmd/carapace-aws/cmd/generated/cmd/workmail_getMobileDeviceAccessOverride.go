package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getMobileDeviceAccessOverrideCmd = &cobra.Command{
	Use:   "get-mobile-device-access-override",
	Short: "Gets the mobile device access override for the given WorkMail organization, user, and device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getMobileDeviceAccessOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_getMobileDeviceAccessOverrideCmd).Standalone()

		workmail_getMobileDeviceAccessOverrideCmd.Flags().String("device-id", "", "The mobile device to which the override applies.")
		workmail_getMobileDeviceAccessOverrideCmd.Flags().String("organization-id", "", "The WorkMail organization to which you want to apply the override.")
		workmail_getMobileDeviceAccessOverrideCmd.Flags().String("user-id", "", "Identifies the WorkMail user for the override.")
		workmail_getMobileDeviceAccessOverrideCmd.MarkFlagRequired("device-id")
		workmail_getMobileDeviceAccessOverrideCmd.MarkFlagRequired("organization-id")
		workmail_getMobileDeviceAccessOverrideCmd.MarkFlagRequired("user-id")
	})
	workmailCmd.AddCommand(workmail_getMobileDeviceAccessOverrideCmd)
}
