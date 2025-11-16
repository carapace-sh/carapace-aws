package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listContactChannelsCmd = &cobra.Command{
	Use:   "list-contact-channels",
	Short: "Lists all contact channels for the specified contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listContactChannelsCmd).Standalone()

	ssmContacts_listContactChannelsCmd.Flags().String("contact-id", "", "The Amazon Resource Name (ARN) of the contact.")
	ssmContacts_listContactChannelsCmd.Flags().String("max-results", "", "The maximum number of contact channels per page.")
	ssmContacts_listContactChannelsCmd.Flags().String("next-token", "", "The pagination token to continue to the next page of results.")
	ssmContacts_listContactChannelsCmd.MarkFlagRequired("contact-id")
	ssmContactsCmd.AddCommand(ssmContacts_listContactChannelsCmd)
}
