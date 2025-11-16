package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listContactListsCmd = &cobra.Command{
	Use:   "list-contact-lists",
	Short: "Lists all of the contact lists available.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listContactListsCmd).Standalone()

	sesv2_listContactListsCmd.Flags().String("next-token", "", "A string token indicating that there might be additional contact lists available to be listed.")
	sesv2_listContactListsCmd.Flags().String("page-size", "", "Maximum number of contact lists to return at once.")
	sesv2Cmd.AddCommand(sesv2_listContactListsCmd)
}
