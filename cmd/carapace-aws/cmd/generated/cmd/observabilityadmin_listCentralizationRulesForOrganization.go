package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_listCentralizationRulesForOrganizationCmd = &cobra.Command{
	Use:   "list-centralization-rules-for-organization",
	Short: "Lists all centralization rules in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_listCentralizationRulesForOrganizationCmd).Standalone()

	observabilityadmin_listCentralizationRulesForOrganizationCmd.Flags().Bool("all-regions", false, "A flag determining whether to return organization centralization rules from all regions or only the current region.")
	observabilityadmin_listCentralizationRulesForOrganizationCmd.Flags().String("max-results", "", "The maximum number of organization centralization rules to return in a single call.")
	observabilityadmin_listCentralizationRulesForOrganizationCmd.Flags().String("next-token", "", "The token for the next set of results.")
	observabilityadmin_listCentralizationRulesForOrganizationCmd.Flags().Bool("no-all-regions", false, "A flag determining whether to return organization centralization rules from all regions or only the current region.")
	observabilityadmin_listCentralizationRulesForOrganizationCmd.Flags().String("rule-name-prefix", "", "A string to filter organization centralization rules whose names begin with the specified prefix.")
	observabilityadmin_listCentralizationRulesForOrganizationCmd.Flag("no-all-regions").Hidden = true
	observabilityadminCmd.AddCommand(observabilityadmin_listCentralizationRulesForOrganizationCmd)
}
