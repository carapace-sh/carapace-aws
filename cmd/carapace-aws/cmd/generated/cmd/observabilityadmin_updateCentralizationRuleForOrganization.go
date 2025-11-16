package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_updateCentralizationRuleForOrganizationCmd = &cobra.Command{
	Use:   "update-centralization-rule-for-organization",
	Short: "Updates an existing centralization rule that applies across an Amazon Web Services Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_updateCentralizationRuleForOrganizationCmd).Standalone()

	observabilityadmin_updateCentralizationRuleForOrganizationCmd.Flags().String("rule", "", "The configuration details for the organization-wide centralization rule, including the source configuration and the destination configuration to centralize telemetry data across the organization.")
	observabilityadmin_updateCentralizationRuleForOrganizationCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the organization centralization rule to update.")
	observabilityadmin_updateCentralizationRuleForOrganizationCmd.MarkFlagRequired("rule")
	observabilityadmin_updateCentralizationRuleForOrganizationCmd.MarkFlagRequired("rule-identifier")
	observabilityadminCmd.AddCommand(observabilityadmin_updateCentralizationRuleForOrganizationCmd)
}
