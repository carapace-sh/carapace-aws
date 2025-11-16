package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listContactsCmd = &cobra.Command{
	Use:   "list-contacts",
	Short: "Lists the contacts present in a specific contact list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listContactsCmd).Standalone()

	sesv2_listContactsCmd.Flags().String("contact-list-name", "", "The name of the contact list.")
	sesv2_listContactsCmd.Flags().String("filter", "", "A filter that can be applied to a list of contacts.")
	sesv2_listContactsCmd.Flags().String("next-token", "", "A string token indicating that there might be additional contacts available to be listed.")
	sesv2_listContactsCmd.Flags().String("page-size", "", "The number of contacts that may be returned at once, which is dependent on if there are more or less contacts than the value of the PageSize.")
	sesv2_listContactsCmd.MarkFlagRequired("contact-list-name")
	sesv2Cmd.AddCommand(sesv2_listContactsCmd)
}
