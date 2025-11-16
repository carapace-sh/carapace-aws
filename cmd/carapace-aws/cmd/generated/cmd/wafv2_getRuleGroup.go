package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getRuleGroupCmd = &cobra.Command{
	Use:   "get-rule-group",
	Short: "Retrieves the specified [RuleGroup]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getRuleGroupCmd).Standalone()

	wafv2_getRuleGroupCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the entity.")
	wafv2_getRuleGroupCmd.Flags().String("id", "", "A unique identifier for the rule group.")
	wafv2_getRuleGroupCmd.Flags().String("name", "", "The name of the rule group.")
	wafv2_getRuleGroupCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2Cmd.AddCommand(wafv2_getRuleGroupCmd)
}
