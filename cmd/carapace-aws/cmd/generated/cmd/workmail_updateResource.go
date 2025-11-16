package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateResourceCmd = &cobra.Command{
	Use:   "update-resource",
	Short: "Updates data for the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_updateResourceCmd).Standalone()

		workmail_updateResourceCmd.Flags().String("booking-options", "", "The resource's booking options to be updated.")
		workmail_updateResourceCmd.Flags().String("description", "", "Updates the resource description.")
		workmail_updateResourceCmd.Flags().String("hidden-from-global-address-list", "", "If enabled, the resource is hidden from the global address list.")
		workmail_updateResourceCmd.Flags().String("name", "", "The name of the resource to be updated.")
		workmail_updateResourceCmd.Flags().String("organization-id", "", "The identifier associated with the organization for which the resource is updated.")
		workmail_updateResourceCmd.Flags().String("resource-id", "", "The identifier of the resource to be updated.")
		workmail_updateResourceCmd.Flags().String("type", "", "Updates the resource type.")
		workmail_updateResourceCmd.MarkFlagRequired("organization-id")
		workmail_updateResourceCmd.MarkFlagRequired("resource-id")
	})
	workmailCmd.AddCommand(workmail_updateResourceCmd)
}
