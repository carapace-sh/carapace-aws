package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_deleteCentralizationRuleForOrganizationCmd = &cobra.Command{
	Use:   "delete-centralization-rule-for-organization",
	Short: "Deletes an organization-wide centralization rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_deleteCentralizationRuleForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_deleteCentralizationRuleForOrganizationCmd).Standalone()

		observabilityadmin_deleteCentralizationRuleForOrganizationCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the organization centralization rule to delete.")
		observabilityadmin_deleteCentralizationRuleForOrganizationCmd.MarkFlagRequired("rule-identifier")
	})
	observabilityadminCmd.AddCommand(observabilityadmin_deleteCentralizationRuleForOrganizationCmd)
}
