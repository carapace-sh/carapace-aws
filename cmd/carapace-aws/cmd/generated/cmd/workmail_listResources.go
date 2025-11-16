package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listResourcesCmd = &cobra.Command{
	Use:   "list-resources",
	Short: "Returns summaries of the organization's resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_listResourcesCmd).Standalone()

		workmail_listResourcesCmd.Flags().String("filters", "", "Limit the resource search results based on the filter criteria.")
		workmail_listResourcesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		workmail_listResourcesCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		workmail_listResourcesCmd.Flags().String("organization-id", "", "The identifier for the organization under which the resources exist.")
		workmail_listResourcesCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_listResourcesCmd)
}
