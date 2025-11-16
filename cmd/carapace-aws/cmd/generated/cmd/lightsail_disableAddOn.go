package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_disableAddOnCmd = &cobra.Command{
	Use:   "disable-add-on",
	Short: "Disables an add-on for an Amazon Lightsail resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_disableAddOnCmd).Standalone()

	lightsail_disableAddOnCmd.Flags().String("add-on-type", "", "The add-on type to disable.")
	lightsail_disableAddOnCmd.Flags().String("resource-name", "", "The name of the source resource for which to disable the add-on.")
	lightsail_disableAddOnCmd.MarkFlagRequired("add-on-type")
	lightsail_disableAddOnCmd.MarkFlagRequired("resource-name")
	lightsailCmd.AddCommand(lightsail_disableAddOnCmd)
}
