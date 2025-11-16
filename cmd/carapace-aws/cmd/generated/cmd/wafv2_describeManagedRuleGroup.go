package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_describeManagedRuleGroupCmd = &cobra.Command{
	Use:   "describe-managed-rule-group",
	Short: "Provides high-level information for a managed rule group, including descriptions of the rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_describeManagedRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_describeManagedRuleGroupCmd).Standalone()

		wafv2_describeManagedRuleGroupCmd.Flags().String("name", "", "The name of the managed rule group.")
		wafv2_describeManagedRuleGroupCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_describeManagedRuleGroupCmd.Flags().String("vendor-name", "", "The name of the managed rule group vendor.")
		wafv2_describeManagedRuleGroupCmd.Flags().String("version-name", "", "The version of the rule group.")
		wafv2_describeManagedRuleGroupCmd.MarkFlagRequired("name")
		wafv2_describeManagedRuleGroupCmd.MarkFlagRequired("scope")
		wafv2_describeManagedRuleGroupCmd.MarkFlagRequired("vendor-name")
	})
	wafv2Cmd.AddCommand(wafv2_describeManagedRuleGroupCmd)
}
