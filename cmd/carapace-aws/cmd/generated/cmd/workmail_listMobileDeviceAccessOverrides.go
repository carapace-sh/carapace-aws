package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listMobileDeviceAccessOverridesCmd = &cobra.Command{
	Use:   "list-mobile-device-access-overrides",
	Short: "Lists all the mobile device access overrides for any given combination of WorkMail organization, user, or device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listMobileDeviceAccessOverridesCmd).Standalone()

	workmail_listMobileDeviceAccessOverridesCmd.Flags().String("device-id", "", "The mobile device to which the access override applies.")
	workmail_listMobileDeviceAccessOverridesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	workmail_listMobileDeviceAccessOverridesCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	workmail_listMobileDeviceAccessOverridesCmd.Flags().String("organization-id", "", "The WorkMail organization under which to list mobile device access overrides.")
	workmail_listMobileDeviceAccessOverridesCmd.Flags().String("user-id", "", "The WorkMail user under which you list the mobile device access overrides.")
	workmail_listMobileDeviceAccessOverridesCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listMobileDeviceAccessOverridesCmd)
}
