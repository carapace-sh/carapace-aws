package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listPersonalAccessTokensCmd = &cobra.Command{
	Use:   "list-personal-access-tokens",
	Short: "Returns a summary of your Personal Access Tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listPersonalAccessTokensCmd).Standalone()

	workmail_listPersonalAccessTokensCmd.Flags().String("max-results", "", "The maximum amount of items that should be returned in a response.")
	workmail_listPersonalAccessTokensCmd.Flags().String("next-token", "", "The token from the previous response to query the next page.")
	workmail_listPersonalAccessTokensCmd.Flags().String("organization-id", "", "The Organization ID.")
	workmail_listPersonalAccessTokensCmd.Flags().String("user-id", "", "The WorkMail User ID.")
	workmail_listPersonalAccessTokensCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listPersonalAccessTokensCmd)
}
