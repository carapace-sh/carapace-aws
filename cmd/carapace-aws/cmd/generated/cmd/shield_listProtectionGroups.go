package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_listProtectionGroupsCmd = &cobra.Command{
	Use:   "list-protection-groups",
	Short: "Retrieves [ProtectionGroup]() objects for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_listProtectionGroupsCmd).Standalone()

	shield_listProtectionGroupsCmd.Flags().String("inclusion-filters", "", "Narrows the set of protection groups that the call retrieves.")
	shield_listProtectionGroupsCmd.Flags().String("max-results", "", "The greatest number of objects that you want Shield Advanced to return to the list request.")
	shield_listProtectionGroupsCmd.Flags().String("next-token", "", "When you request a list of objects from Shield Advanced, if the response does not include all of the remaining available objects, Shield Advanced includes a `NextToken` value in the response.")
	shieldCmd.AddCommand(shield_listProtectionGroupsCmd)
}
