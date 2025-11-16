package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getPersonalAccessTokenMetadataCmd = &cobra.Command{
	Use:   "get-personal-access-token-metadata",
	Short: "Requests details of a specific Personal Access Token within the WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getPersonalAccessTokenMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_getPersonalAccessTokenMetadataCmd).Standalone()

		workmail_getPersonalAccessTokenMetadataCmd.Flags().String("organization-id", "", "The Organization ID.")
		workmail_getPersonalAccessTokenMetadataCmd.Flags().String("personal-access-token-id", "", "The Personal Access Token ID.")
		workmail_getPersonalAccessTokenMetadataCmd.MarkFlagRequired("organization-id")
		workmail_getPersonalAccessTokenMetadataCmd.MarkFlagRequired("personal-access-token-id")
	})
	workmailCmd.AddCommand(workmail_getPersonalAccessTokenMetadataCmd)
}
