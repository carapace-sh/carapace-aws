package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Deletes the specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_deleteRuleCmd).Standalone()

	elbv2_deleteRuleCmd.Flags().String("rule-arn", "", "The Amazon Resource Name (ARN) of the rule.")
	elbv2_deleteRuleCmd.MarkFlagRequired("rule-arn")
	elbv2Cmd.AddCommand(elbv2_deleteRuleCmd)
}
