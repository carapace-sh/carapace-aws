package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_disassociateDelegateFromResourceCmd = &cobra.Command{
	Use:   "disassociate-delegate-from-resource",
	Short: "Removes a member from the resource's set of delegates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_disassociateDelegateFromResourceCmd).Standalone()

	workmail_disassociateDelegateFromResourceCmd.Flags().String("entity-id", "", "The identifier for the member (user, group) to be removed from the resource's delegates.")
	workmail_disassociateDelegateFromResourceCmd.Flags().String("organization-id", "", "The identifier for the organization under which the resource exists.")
	workmail_disassociateDelegateFromResourceCmd.Flags().String("resource-id", "", "The identifier of the resource from which delegates' set members are removed.")
	workmail_disassociateDelegateFromResourceCmd.MarkFlagRequired("entity-id")
	workmail_disassociateDelegateFromResourceCmd.MarkFlagRequired("organization-id")
	workmail_disassociateDelegateFromResourceCmd.MarkFlagRequired("resource-id")
	workmailCmd.AddCommand(workmail_disassociateDelegateFromResourceCmd)
}
