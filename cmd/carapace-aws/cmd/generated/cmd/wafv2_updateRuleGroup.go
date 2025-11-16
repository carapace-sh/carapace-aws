package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_updateRuleGroupCmd = &cobra.Command{
	Use:   "update-rule-group",
	Short: "Updates the specified [RuleGroup]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_updateRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_updateRuleGroupCmd).Standalone()

		wafv2_updateRuleGroupCmd.Flags().String("custom-response-bodies", "", "A map of custom response keys and content bodies.")
		wafv2_updateRuleGroupCmd.Flags().String("description", "", "A description of the rule group that helps with identification.")
		wafv2_updateRuleGroupCmd.Flags().String("id", "", "A unique identifier for the rule group.")
		wafv2_updateRuleGroupCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
		wafv2_updateRuleGroupCmd.Flags().String("name", "", "The name of the rule group.")
		wafv2_updateRuleGroupCmd.Flags().String("rules", "", "The [Rule]() statements used to identify the web requests that you want to manage.")
		wafv2_updateRuleGroupCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_updateRuleGroupCmd.Flags().String("visibility-config", "", "Defines and enables Amazon CloudWatch metrics and web request sample collection.")
		wafv2_updateRuleGroupCmd.MarkFlagRequired("id")
		wafv2_updateRuleGroupCmd.MarkFlagRequired("lock-token")
		wafv2_updateRuleGroupCmd.MarkFlagRequired("name")
		wafv2_updateRuleGroupCmd.MarkFlagRequired("scope")
		wafv2_updateRuleGroupCmd.MarkFlagRequired("visibility-config")
	})
	wafv2Cmd.AddCommand(wafv2_updateRuleGroupCmd)
}
