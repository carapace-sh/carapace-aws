package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getDefaultRetentionPolicyCmd = &cobra.Command{
	Use:   "get-default-retention-policy",
	Short: "Gets the default retention policy details for the specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getDefaultRetentionPolicyCmd).Standalone()

	workmail_getDefaultRetentionPolicyCmd.Flags().String("organization-id", "", "The organization ID.")
	workmail_getDefaultRetentionPolicyCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_getDefaultRetentionPolicyCmd)
}
