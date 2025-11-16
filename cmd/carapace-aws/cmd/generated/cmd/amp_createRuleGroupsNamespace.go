package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_createRuleGroupsNamespaceCmd = &cobra.Command{
	Use:   "create-rule-groups-namespace",
	Short: "The `CreateRuleGroupsNamespace` operation creates a rule groups namespace within a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_createRuleGroupsNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_createRuleGroupsNamespaceCmd).Standalone()

		amp_createRuleGroupsNamespaceCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_createRuleGroupsNamespaceCmd.Flags().String("data", "", "The rules file to use in the new namespace.")
		amp_createRuleGroupsNamespaceCmd.Flags().String("name", "", "The name for the new rule groups namespace.")
		amp_createRuleGroupsNamespaceCmd.Flags().String("tags", "", "The list of tag keys and values to associate with the rule groups namespace.")
		amp_createRuleGroupsNamespaceCmd.Flags().String("workspace-id", "", "The ID of the workspace to add the rule groups namespace.")
		amp_createRuleGroupsNamespaceCmd.MarkFlagRequired("data")
		amp_createRuleGroupsNamespaceCmd.MarkFlagRequired("name")
		amp_createRuleGroupsNamespaceCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_createRuleGroupsNamespaceCmd)
}
