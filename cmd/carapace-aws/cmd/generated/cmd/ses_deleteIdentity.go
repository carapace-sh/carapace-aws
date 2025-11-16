package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteIdentityCmd = &cobra.Command{
	Use:   "delete-identity",
	Short: "Deletes the specified identity (an email address or a domain) from the list of verified identities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteIdentityCmd).Standalone()

	ses_deleteIdentityCmd.Flags().String("identity", "", "The identity to be removed from the list of identities for the Amazon Web Services account.")
	ses_deleteIdentityCmd.MarkFlagRequired("identity")
	sesCmd.AddCommand(ses_deleteIdentityCmd)
}
