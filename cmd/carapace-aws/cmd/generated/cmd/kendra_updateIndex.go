package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateIndexCmd = &cobra.Command{
	Use:   "update-index",
	Short: "Updates an Amazon Kendra index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateIndexCmd).Standalone()

	kendra_updateIndexCmd.Flags().String("capacity-units", "", "Sets the number of additional document storage and query capacity units that should be used by the index.")
	kendra_updateIndexCmd.Flags().String("description", "", "A new description for the index.")
	kendra_updateIndexCmd.Flags().String("document-metadata-configuration-updates", "", "The document metadata configuration you want to update for the index.")
	kendra_updateIndexCmd.Flags().String("id", "", "The identifier of the index you want to update.")
	kendra_updateIndexCmd.Flags().String("name", "", "A new name for the index.")
	kendra_updateIndexCmd.Flags().String("role-arn", "", "An Identity and Access Management (IAM) role that gives Amazon Kendra permission to access Amazon CloudWatch logs and metrics.")
	kendra_updateIndexCmd.Flags().String("user-context-policy", "", "The user context policy.")
	kendra_updateIndexCmd.Flags().String("user-group-resolution-configuration", "", "Gets users and groups from IAM Identity Center identity source.")
	kendra_updateIndexCmd.Flags().String("user-token-configurations", "", "The user token configuration.")
	kendra_updateIndexCmd.MarkFlagRequired("id")
	kendraCmd.AddCommand(kendra_updateIndexCmd)
}
