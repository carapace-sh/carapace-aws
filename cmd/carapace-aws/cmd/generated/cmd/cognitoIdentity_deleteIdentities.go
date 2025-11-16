package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_deleteIdentitiesCmd = &cobra.Command{
	Use:   "delete-identities",
	Short: "Deletes identities from an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_deleteIdentitiesCmd).Standalone()

	cognitoIdentity_deleteIdentitiesCmd.Flags().String("identity-ids-to-delete", "", "A list of 1-60 identities that you want to delete.")
	cognitoIdentity_deleteIdentitiesCmd.MarkFlagRequired("identity-ids-to-delete")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_deleteIdentitiesCmd)
}
