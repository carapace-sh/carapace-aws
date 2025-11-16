package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listIdentitiesCmd = &cobra.Command{
	Use:   "list-identities",
	Short: "Returns a list containing all of the identities (email addresses and domains) for your Amazon Web Services account in the current Amazon Web Services Region, regardless of verification status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listIdentitiesCmd).Standalone()

	ses_listIdentitiesCmd.Flags().String("identity-type", "", "The type of the identities to list.")
	ses_listIdentitiesCmd.Flags().String("max-items", "", "The maximum number of identities per page.")
	ses_listIdentitiesCmd.Flags().String("next-token", "", "The token to use for pagination.")
	sesCmd.AddCommand(ses_listIdentitiesCmd)
}
