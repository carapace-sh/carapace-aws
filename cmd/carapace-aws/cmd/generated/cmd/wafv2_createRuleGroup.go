package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_createRuleGroupCmd = &cobra.Command{
	Use:   "create-rule-group",
	Short: "Creates a [RuleGroup]() per the specifications provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_createRuleGroupCmd).Standalone()

	wafv2_createRuleGroupCmd.Flags().String("capacity", "", "The web ACL capacity units (WCUs) required for this rule group.")
	wafv2_createRuleGroupCmd.Flags().String("custom-response-bodies", "", "A map of custom response keys and content bodies.")
	wafv2_createRuleGroupCmd.Flags().String("description", "", "A description of the rule group that helps with identification.")
	wafv2_createRuleGroupCmd.Flags().String("name", "", "The name of the rule group.")
	wafv2_createRuleGroupCmd.Flags().String("rules", "", "The [Rule]() statements used to identify the web requests that you want to manage.")
	wafv2_createRuleGroupCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_createRuleGroupCmd.Flags().String("tags", "", "An array of key:value pairs to associate with the resource.")
	wafv2_createRuleGroupCmd.Flags().String("visibility-config", "", "Defines and enables Amazon CloudWatch metrics and web request sample collection.")
	wafv2_createRuleGroupCmd.MarkFlagRequired("capacity")
	wafv2_createRuleGroupCmd.MarkFlagRequired("name")
	wafv2_createRuleGroupCmd.MarkFlagRequired("scope")
	wafv2_createRuleGroupCmd.MarkFlagRequired("visibility-config")
	wafv2Cmd.AddCommand(wafv2_createRuleGroupCmd)
}
