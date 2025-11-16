package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "Creates a rule for the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_createRuleCmd).Standalone()

	elbv2_createRuleCmd.Flags().String("actions", "", "The actions.")
	elbv2_createRuleCmd.Flags().String("conditions", "", "The conditions.")
	elbv2_createRuleCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
	elbv2_createRuleCmd.Flags().String("priority", "", "The rule priority.")
	elbv2_createRuleCmd.Flags().String("tags", "", "The tags to assign to the rule.")
	elbv2_createRuleCmd.Flags().String("transforms", "", "The transforms to apply to requests that match this rule.")
	elbv2_createRuleCmd.MarkFlagRequired("actions")
	elbv2_createRuleCmd.MarkFlagRequired("conditions")
	elbv2_createRuleCmd.MarkFlagRequired("listener-arn")
	elbv2_createRuleCmd.MarkFlagRequired("priority")
	elbv2Cmd.AddCommand(elbv2_createRuleCmd)
}
