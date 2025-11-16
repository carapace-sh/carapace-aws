package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listEmailIdentitiesCmd = &cobra.Command{
	Use:   "list-email-identities",
	Short: "Returns a list of all of the email identities that are associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listEmailIdentitiesCmd).Standalone()

	sesv2_listEmailIdentitiesCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListEmailIdentities` to indicate the position in the list of identities.")
	sesv2_listEmailIdentitiesCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListEmailIdentities`.")
	sesv2Cmd.AddCommand(sesv2_listEmailIdentitiesCmd)
}
