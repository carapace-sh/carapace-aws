package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listAssociatedContactsCmd = &cobra.Command{
	Use:   "list-associated-contacts",
	Short: "Provides information about contact tree, a list of associated contacts with a unique identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listAssociatedContactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listAssociatedContactsCmd).Standalone()

		connect_listAssociatedContactsCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_listAssociatedContactsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listAssociatedContactsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listAssociatedContactsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listAssociatedContactsCmd.MarkFlagRequired("contact-id")
		connect_listAssociatedContactsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listAssociatedContactsCmd)
}
