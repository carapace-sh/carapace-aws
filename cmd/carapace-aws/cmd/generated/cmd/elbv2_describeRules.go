package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeRulesCmd = &cobra.Command{
	Use:   "describe-rules",
	Short: "Describes the specified rules or the rules for the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeRulesCmd).Standalone()

		elbv2_describeRulesCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
		elbv2_describeRulesCmd.Flags().String("marker", "", "The marker for the next set of results.")
		elbv2_describeRulesCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
		elbv2_describeRulesCmd.Flags().String("rule-arns", "", "The Amazon Resource Names (ARN) of the rules.")
	})
	elbv2Cmd.AddCommand(elbv2_describeRulesCmd)
}
