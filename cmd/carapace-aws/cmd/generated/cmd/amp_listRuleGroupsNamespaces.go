package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_listRuleGroupsNamespacesCmd = &cobra.Command{
	Use:   "list-rule-groups-namespaces",
	Short: "Returns a list of rule groups namespaces in a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_listRuleGroupsNamespacesCmd).Standalone()

	amp_listRuleGroupsNamespacesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	amp_listRuleGroupsNamespacesCmd.Flags().String("name", "", "Use this parameter to filter the rule groups namespaces that are returned.")
	amp_listRuleGroupsNamespacesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	amp_listRuleGroupsNamespacesCmd.Flags().String("workspace-id", "", "The ID of the workspace containing the rule groups namespaces.")
	amp_listRuleGroupsNamespacesCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_listRuleGroupsNamespacesCmd)
}
