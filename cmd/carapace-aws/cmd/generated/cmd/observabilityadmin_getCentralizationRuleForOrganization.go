package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_getCentralizationRuleForOrganizationCmd = &cobra.Command{
	Use:   "get-centralization-rule-for-organization",
	Short: "Retrieves the details of a specific organization centralization rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_getCentralizationRuleForOrganizationCmd).Standalone()

	observabilityadmin_getCentralizationRuleForOrganizationCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the organization centralization rule to retrieve.")
	observabilityadmin_getCentralizationRuleForOrganizationCmd.MarkFlagRequired("rule-identifier")
	observabilityadminCmd.AddCommand(observabilityadmin_getCentralizationRuleForOrganizationCmd)
}
