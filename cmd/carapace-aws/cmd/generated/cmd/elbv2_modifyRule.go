package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyRuleCmd = &cobra.Command{
	Use:   "modify-rule",
	Short: "Replaces the specified properties of the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyRuleCmd).Standalone()

	elbv2_modifyRuleCmd.Flags().String("actions", "", "The actions.")
	elbv2_modifyRuleCmd.Flags().String("conditions", "", "The conditions.")
	elbv2_modifyRuleCmd.Flags().String("reset-transforms", "", "Indicates whether to remove all transforms from the rule.")
	elbv2_modifyRuleCmd.Flags().String("rule-arn", "", "The Amazon Resource Name (ARN) of the rule.")
	elbv2_modifyRuleCmd.Flags().String("transforms", "", "The transforms to apply to requests that match this rule.")
	elbv2_modifyRuleCmd.MarkFlagRequired("rule-arn")
	elbv2Cmd.AddCommand(elbv2_modifyRuleCmd)
}
