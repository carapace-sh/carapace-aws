package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeOrganizationConformancePackStatusesCmd = &cobra.Command{
	Use:   "describe-organization-conformance-pack-statuses",
	Short: "Provides organization conformance pack deployment status for an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeOrganizationConformancePackStatusesCmd).Standalone()

	config_describeOrganizationConformancePackStatusesCmd.Flags().String("limit", "", "The maximum number of OrganizationConformancePackStatuses returned on each page.")
	config_describeOrganizationConformancePackStatusesCmd.Flags().String("next-token", "", "The nextToken string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_describeOrganizationConformancePackStatusesCmd.Flags().String("organization-conformance-pack-names", "", "The names of organization conformance packs for which you want status details.")
	configCmd.AddCommand(config_describeOrganizationConformancePackStatusesCmd)
}
