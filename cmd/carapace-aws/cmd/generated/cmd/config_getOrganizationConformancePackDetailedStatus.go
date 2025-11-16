package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getOrganizationConformancePackDetailedStatusCmd = &cobra.Command{
	Use:   "get-organization-conformance-pack-detailed-status",
	Short: "Returns detailed status for each member account within an organization for a given organization conformance pack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getOrganizationConformancePackDetailedStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getOrganizationConformancePackDetailedStatusCmd).Standalone()

		config_getOrganizationConformancePackDetailedStatusCmd.Flags().String("filters", "", "An `OrganizationResourceDetailedStatusFilters` object.")
		config_getOrganizationConformancePackDetailedStatusCmd.Flags().String("limit", "", "The maximum number of `OrganizationConformancePackDetailedStatuses` returned on each page.")
		config_getOrganizationConformancePackDetailedStatusCmd.Flags().String("next-token", "", "The nextToken string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_getOrganizationConformancePackDetailedStatusCmd.Flags().String("organization-conformance-pack-name", "", "The name of organization conformance pack for which you want status details for member accounts.")
		config_getOrganizationConformancePackDetailedStatusCmd.MarkFlagRequired("organization-conformance-pack-name")
	})
	configCmd.AddCommand(config_getOrganizationConformancePackDetailedStatusCmd)
}
