package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeOrganizationConformancePacksCmd = &cobra.Command{
	Use:   "describe-organization-conformance-packs",
	Short: "Returns a list of organization conformance packs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeOrganizationConformancePacksCmd).Standalone()

	config_describeOrganizationConformancePacksCmd.Flags().String("limit", "", "The maximum number of organization config packs returned on each page.")
	config_describeOrganizationConformancePacksCmd.Flags().String("next-token", "", "The nextToken string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_describeOrganizationConformancePacksCmd.Flags().String("organization-conformance-pack-names", "", "The name that you assign to an organization conformance pack.")
	configCmd.AddCommand(config_describeOrganizationConformancePacksCmd)
}
