package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteOrganizationConfigRuleCmd = &cobra.Command{
	Use:   "delete-organization-config-rule",
	Short: "Deletes the specified organization Config rule and all of its evaluation results from all member accounts in that organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteOrganizationConfigRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteOrganizationConfigRuleCmd).Standalone()

		config_deleteOrganizationConfigRuleCmd.Flags().String("organization-config-rule-name", "", "The name of organization Config rule that you want to delete.")
		config_deleteOrganizationConfigRuleCmd.MarkFlagRequired("organization-config-rule-name")
	})
	configCmd.AddCommand(config_deleteOrganizationConfigRuleCmd)
}
