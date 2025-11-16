package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createAliasCmd = &cobra.Command{
	Use:   "create-alias",
	Short: "Adds an alias to the set of a given member (user or group) of WorkMail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_createAliasCmd).Standalone()

		workmail_createAliasCmd.Flags().String("alias", "", "The alias to add to the member set.")
		workmail_createAliasCmd.Flags().String("entity-id", "", "The member (user or group) to which this alias is added.")
		workmail_createAliasCmd.Flags().String("organization-id", "", "The organization under which the member (user or group) exists.")
		workmail_createAliasCmd.MarkFlagRequired("alias")
		workmail_createAliasCmd.MarkFlagRequired("entity-id")
		workmail_createAliasCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_createAliasCmd)
}
