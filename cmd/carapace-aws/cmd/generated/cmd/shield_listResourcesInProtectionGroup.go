package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_listResourcesInProtectionGroupCmd = &cobra.Command{
	Use:   "list-resources-in-protection-group",
	Short: "Retrieves the resources that are included in the protection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_listResourcesInProtectionGroupCmd).Standalone()

	shield_listResourcesInProtectionGroupCmd.Flags().String("max-results", "", "The greatest number of objects that you want Shield Advanced to return to the list request.")
	shield_listResourcesInProtectionGroupCmd.Flags().String("next-token", "", "When you request a list of objects from Shield Advanced, if the response does not include all of the remaining available objects, Shield Advanced includes a `NextToken` value in the response.")
	shield_listResourcesInProtectionGroupCmd.Flags().String("protection-group-id", "", "The name of the protection group.")
	shield_listResourcesInProtectionGroupCmd.MarkFlagRequired("protection-group-id")
	shieldCmd.AddCommand(shield_listResourcesInProtectionGroupCmd)
}
