package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deleteRuleGroupCmd = &cobra.Command{
	Use:   "delete-rule-group",
	Short: "Deletes the specified [RuleGroup]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deleteRuleGroupCmd).Standalone()

	wafv2_deleteRuleGroupCmd.Flags().String("id", "", "A unique identifier for the rule group.")
	wafv2_deleteRuleGroupCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
	wafv2_deleteRuleGroupCmd.Flags().String("name", "", "The name of the rule group.")
	wafv2_deleteRuleGroupCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_deleteRuleGroupCmd.MarkFlagRequired("id")
	wafv2_deleteRuleGroupCmd.MarkFlagRequired("lock-token")
	wafv2_deleteRuleGroupCmd.MarkFlagRequired("name")
	wafv2_deleteRuleGroupCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_deleteRuleGroupCmd)
}
