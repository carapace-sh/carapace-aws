package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deletePersonalAccessTokenCmd = &cobra.Command{
	Use:   "delete-personal-access-token",
	Short: "Deletes the Personal Access Token from the provided WorkMail Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deletePersonalAccessTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_deletePersonalAccessTokenCmd).Standalone()

		workmail_deletePersonalAccessTokenCmd.Flags().String("organization-id", "", "The Organization ID.")
		workmail_deletePersonalAccessTokenCmd.Flags().String("personal-access-token-id", "", "The Personal Access Token ID.")
		workmail_deletePersonalAccessTokenCmd.MarkFlagRequired("organization-id")
		workmail_deletePersonalAccessTokenCmd.MarkFlagRequired("personal-access-token-id")
	})
	workmailCmd.AddCommand(workmail_deletePersonalAccessTokenCmd)
}
