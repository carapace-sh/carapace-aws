package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listOrganizationsCmd = &cobra.Command{
	Use:   "list-organizations",
	Short: "Returns summaries of the customer's organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listOrganizationsCmd).Standalone()

	workmail_listOrganizationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	workmail_listOrganizationsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	workmailCmd.AddCommand(workmail_listOrganizationsCmd)
}
