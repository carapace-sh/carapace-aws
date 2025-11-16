package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_associateDelegateToResourceCmd = &cobra.Command{
	Use:   "associate-delegate-to-resource",
	Short: "Adds a member (user or group) to the resource's set of delegates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_associateDelegateToResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_associateDelegateToResourceCmd).Standalone()

		workmail_associateDelegateToResourceCmd.Flags().String("entity-id", "", "The member (user or group) to associate to the resource.")
		workmail_associateDelegateToResourceCmd.Flags().String("organization-id", "", "The organization under which the resource exists.")
		workmail_associateDelegateToResourceCmd.Flags().String("resource-id", "", "The resource for which members (users or groups) are associated.")
		workmail_associateDelegateToResourceCmd.MarkFlagRequired("entity-id")
		workmail_associateDelegateToResourceCmd.MarkFlagRequired("organization-id")
		workmail_associateDelegateToResourceCmd.MarkFlagRequired("resource-id")
	})
	workmailCmd.AddCommand(workmail_associateDelegateToResourceCmd)
}
