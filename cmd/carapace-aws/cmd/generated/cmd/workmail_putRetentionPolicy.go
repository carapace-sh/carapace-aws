package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_putRetentionPolicyCmd = &cobra.Command{
	Use:   "put-retention-policy",
	Short: "Puts a retention policy to the specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_putRetentionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_putRetentionPolicyCmd).Standalone()

		workmail_putRetentionPolicyCmd.Flags().String("description", "", "The retention policy description.")
		workmail_putRetentionPolicyCmd.Flags().String("folder-configurations", "", "The retention policy folder configurations.")
		workmail_putRetentionPolicyCmd.Flags().String("id", "", "The retention policy ID.")
		workmail_putRetentionPolicyCmd.Flags().String("name", "", "The retention policy name.")
		workmail_putRetentionPolicyCmd.Flags().String("organization-id", "", "The organization ID.")
		workmail_putRetentionPolicyCmd.MarkFlagRequired("folder-configurations")
		workmail_putRetentionPolicyCmd.MarkFlagRequired("name")
		workmail_putRetentionPolicyCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_putRetentionPolicyCmd)
}
