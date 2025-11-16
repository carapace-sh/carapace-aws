package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listPagesByContactCmd = &cobra.Command{
	Use:   "list-pages-by-contact",
	Short: "Lists the engagements to a contact's contact channels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listPagesByContactCmd).Standalone()

	ssmContacts_listPagesByContactCmd.Flags().String("contact-id", "", "The Amazon Resource Name (ARN) of the contact you are retrieving engagements for.")
	ssmContacts_listPagesByContactCmd.Flags().String("max-results", "", "The maximum number of engagements to contact channels to list per page of results.")
	ssmContacts_listPagesByContactCmd.Flags().String("next-token", "", "The pagination token to continue to the next page of results.")
	ssmContacts_listPagesByContactCmd.MarkFlagRequired("contact-id")
	ssmContactsCmd.AddCommand(ssmContacts_listPagesByContactCmd)
}
