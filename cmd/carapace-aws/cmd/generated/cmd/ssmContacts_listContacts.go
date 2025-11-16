package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listContactsCmd = &cobra.Command{
	Use:   "list-contacts",
	Short: "Lists all contacts and escalation plans in Incident Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listContactsCmd).Standalone()

	ssmContacts_listContactsCmd.Flags().String("alias-prefix", "", "Used to list only contacts who's aliases start with the specified prefix.")
	ssmContacts_listContactsCmd.Flags().String("max-results", "", "The maximum number of contacts and escalation plans per page of results.")
	ssmContacts_listContactsCmd.Flags().String("next-token", "", "The pagination token to continue to the next page of results.")
	ssmContacts_listContactsCmd.Flags().String("type", "", "The type of contact.")
	ssmContactsCmd.AddCommand(ssmContacts_listContactsCmd)
}
