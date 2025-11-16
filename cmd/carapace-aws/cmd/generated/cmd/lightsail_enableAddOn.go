package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_enableAddOnCmd = &cobra.Command{
	Use:   "enable-add-on",
	Short: "Enables or modifies an add-on for an Amazon Lightsail resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_enableAddOnCmd).Standalone()

	lightsail_enableAddOnCmd.Flags().String("add-on-request", "", "An array of strings representing the add-on to enable or modify.")
	lightsail_enableAddOnCmd.Flags().String("resource-name", "", "The name of the source resource for which to enable or modify the add-on.")
	lightsail_enableAddOnCmd.MarkFlagRequired("add-on-request")
	lightsail_enableAddOnCmd.MarkFlagRequired("resource-name")
	lightsailCmd.AddCommand(lightsail_enableAddOnCmd)
}
