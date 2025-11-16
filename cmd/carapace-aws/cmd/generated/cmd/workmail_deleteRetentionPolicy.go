package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteRetentionPolicyCmd = &cobra.Command{
	Use:   "delete-retention-policy",
	Short: "Deletes the specified retention policy from the specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteRetentionPolicyCmd).Standalone()

	workmail_deleteRetentionPolicyCmd.Flags().String("id", "", "The retention policy ID.")
	workmail_deleteRetentionPolicyCmd.Flags().String("organization-id", "", "The organization ID.")
	workmail_deleteRetentionPolicyCmd.MarkFlagRequired("id")
	workmail_deleteRetentionPolicyCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteRetentionPolicyCmd)
}
