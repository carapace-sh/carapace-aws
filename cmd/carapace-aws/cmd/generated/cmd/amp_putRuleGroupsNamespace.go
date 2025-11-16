package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_putRuleGroupsNamespaceCmd = &cobra.Command{
	Use:   "put-rule-groups-namespace",
	Short: "Updates an existing rule groups namespace within a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_putRuleGroupsNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_putRuleGroupsNamespaceCmd).Standalone()

		amp_putRuleGroupsNamespaceCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_putRuleGroupsNamespaceCmd.Flags().String("data", "", "The new rules file to use in the namespace.")
		amp_putRuleGroupsNamespaceCmd.Flags().String("name", "", "The name of the rule groups namespace that you are updating.")
		amp_putRuleGroupsNamespaceCmd.Flags().String("workspace-id", "", "The ID of the workspace where you are updating the rule groups namespace.")
		amp_putRuleGroupsNamespaceCmd.MarkFlagRequired("data")
		amp_putRuleGroupsNamespaceCmd.MarkFlagRequired("name")
		amp_putRuleGroupsNamespaceCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_putRuleGroupsNamespaceCmd)
}
