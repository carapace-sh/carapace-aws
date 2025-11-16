package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_createCentralizationRuleForOrganizationCmd = &cobra.Command{
	Use:   "create-centralization-rule-for-organization",
	Short: "Creates a centralization rule that applies across an Amazon Web Services Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_createCentralizationRuleForOrganizationCmd).Standalone()

	observabilityadmin_createCentralizationRuleForOrganizationCmd.Flags().String("rule", "", "The configuration details for the organization-wide centralization rule, including the source configuration and the destination configuration to centralize telemetry data across the organization.")
	observabilityadmin_createCentralizationRuleForOrganizationCmd.Flags().String("rule-name", "", "A unique name for the organization-wide centralization rule being created.")
	observabilityadmin_createCentralizationRuleForOrganizationCmd.Flags().String("tags", "", "The key-value pairs to associate with the organization telemetry rule resource for categorization and management purposes.")
	observabilityadmin_createCentralizationRuleForOrganizationCmd.MarkFlagRequired("rule")
	observabilityadmin_createCentralizationRuleForOrganizationCmd.MarkFlagRequired("rule-name")
	observabilityadminCmd.AddCommand(observabilityadmin_createCentralizationRuleForOrganizationCmd)
}
