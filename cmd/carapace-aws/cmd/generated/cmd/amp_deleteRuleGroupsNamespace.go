package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteRuleGroupsNamespaceCmd = &cobra.Command{
	Use:   "delete-rule-groups-namespace",
	Short: "Deletes one rule groups namespace and its associated rule groups definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteRuleGroupsNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_deleteRuleGroupsNamespaceCmd).Standalone()

		amp_deleteRuleGroupsNamespaceCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_deleteRuleGroupsNamespaceCmd.Flags().String("name", "", "The name of the rule groups namespace to delete.")
		amp_deleteRuleGroupsNamespaceCmd.Flags().String("workspace-id", "", "The ID of the workspace containing the rule groups namespace and definition to delete.")
		amp_deleteRuleGroupsNamespaceCmd.MarkFlagRequired("name")
		amp_deleteRuleGroupsNamespaceCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_deleteRuleGroupsNamespaceCmd)
}
