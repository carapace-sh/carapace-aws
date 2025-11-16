package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteAliasCmd = &cobra.Command{
	Use:   "delete-alias",
	Short: "Remove one or more specified aliases from a set of aliases for a given user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteAliasCmd).Standalone()

	workmail_deleteAliasCmd.Flags().String("alias", "", "The aliases to be removed from the user's set of aliases.")
	workmail_deleteAliasCmd.Flags().String("entity-id", "", "The identifier for the member (user or group) from which to have the aliases removed.")
	workmail_deleteAliasCmd.Flags().String("organization-id", "", "The identifier for the organization under which the user exists.")
	workmail_deleteAliasCmd.MarkFlagRequired("alias")
	workmail_deleteAliasCmd.MarkFlagRequired("entity-id")
	workmail_deleteAliasCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteAliasCmd)
}
