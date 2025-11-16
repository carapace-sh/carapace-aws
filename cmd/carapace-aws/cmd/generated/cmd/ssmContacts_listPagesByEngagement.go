package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listPagesByEngagementCmd = &cobra.Command{
	Use:   "list-pages-by-engagement",
	Short: "Lists the engagements to contact channels that occurred by engaging a contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listPagesByEngagementCmd).Standalone()

	ssmContacts_listPagesByEngagementCmd.Flags().String("engagement-id", "", "The Amazon Resource Name (ARN) of the engagement.")
	ssmContacts_listPagesByEngagementCmd.Flags().String("max-results", "", "The maximum number of engagements to contact channels to list per page of results.")
	ssmContacts_listPagesByEngagementCmd.Flags().String("next-token", "", "The pagination token to continue to the next page of results.")
	ssmContacts_listPagesByEngagementCmd.MarkFlagRequired("engagement-id")
	ssmContactsCmd.AddCommand(ssmContacts_listPagesByEngagementCmd)
}
