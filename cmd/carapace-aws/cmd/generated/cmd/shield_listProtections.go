package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_listProtectionsCmd = &cobra.Command{
	Use:   "list-protections",
	Short: "Retrieves [Protection]() objects for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_listProtectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_listProtectionsCmd).Standalone()

		shield_listProtectionsCmd.Flags().String("inclusion-filters", "", "Narrows the set of protections that the call retrieves.")
		shield_listProtectionsCmd.Flags().String("max-results", "", "The greatest number of objects that you want Shield Advanced to return to the list request.")
		shield_listProtectionsCmd.Flags().String("next-token", "", "When you request a list of objects from Shield Advanced, if the response does not include all of the remaining available objects, Shield Advanced includes a `NextToken` value in the response.")
	})
	shieldCmd.AddCommand(shield_listProtectionsCmd)
}
