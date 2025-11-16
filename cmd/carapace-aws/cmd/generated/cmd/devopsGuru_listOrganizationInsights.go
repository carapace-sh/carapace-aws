package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listOrganizationInsightsCmd = &cobra.Command{
	Use:   "list-organization-insights",
	Short: "Returns a list of insights associated with the account or OU Id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listOrganizationInsightsCmd).Standalone()

	devopsGuru_listOrganizationInsightsCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account.")
	devopsGuru_listOrganizationInsightsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	devopsGuru_listOrganizationInsightsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuru_listOrganizationInsightsCmd.Flags().String("organizational-unit-ids", "", "The ID of the organizational unit.")
	devopsGuru_listOrganizationInsightsCmd.Flags().String("status-filter", "", "")
	devopsGuru_listOrganizationInsightsCmd.MarkFlagRequired("status-filter")
	devopsGuruCmd.AddCommand(devopsGuru_listOrganizationInsightsCmd)
}
