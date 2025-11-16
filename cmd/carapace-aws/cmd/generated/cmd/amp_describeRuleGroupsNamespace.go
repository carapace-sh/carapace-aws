package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeRuleGroupsNamespaceCmd = &cobra.Command{
	Use:   "describe-rule-groups-namespace",
	Short: "Returns complete information about one rule groups namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeRuleGroupsNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_describeRuleGroupsNamespaceCmd).Standalone()

		amp_describeRuleGroupsNamespaceCmd.Flags().String("name", "", "The name of the rule groups namespace that you want information for.")
		amp_describeRuleGroupsNamespaceCmd.Flags().String("workspace-id", "", "The ID of the workspace containing the rule groups namespace.")
		amp_describeRuleGroupsNamespaceCmd.MarkFlagRequired("name")
		amp_describeRuleGroupsNamespaceCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_describeRuleGroupsNamespaceCmd)
}
