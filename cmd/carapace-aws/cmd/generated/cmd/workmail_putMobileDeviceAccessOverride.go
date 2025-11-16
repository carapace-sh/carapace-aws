package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_putMobileDeviceAccessOverrideCmd = &cobra.Command{
	Use:   "put-mobile-device-access-override",
	Short: "Creates or updates a mobile device access override for the given WorkMail organization, user, and device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_putMobileDeviceAccessOverrideCmd).Standalone()

	workmail_putMobileDeviceAccessOverrideCmd.Flags().String("description", "", "A description of the override.")
	workmail_putMobileDeviceAccessOverrideCmd.Flags().String("device-id", "", "The mobile device for which you create the override.")
	workmail_putMobileDeviceAccessOverrideCmd.Flags().String("effect", "", "The effect of the override, `ALLOW` or `DENY`.")
	workmail_putMobileDeviceAccessOverrideCmd.Flags().String("organization-id", "", "Identifies the WorkMail organization for which you create the override.")
	workmail_putMobileDeviceAccessOverrideCmd.Flags().String("user-id", "", "The WorkMail user for which you create the override.")
	workmail_putMobileDeviceAccessOverrideCmd.MarkFlagRequired("device-id")
	workmail_putMobileDeviceAccessOverrideCmd.MarkFlagRequired("effect")
	workmail_putMobileDeviceAccessOverrideCmd.MarkFlagRequired("organization-id")
	workmail_putMobileDeviceAccessOverrideCmd.MarkFlagRequired("user-id")
	workmailCmd.AddCommand(workmail_putMobileDeviceAccessOverrideCmd)
}
