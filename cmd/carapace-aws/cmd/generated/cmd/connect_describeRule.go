package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeRuleCmd = &cobra.Command{
	Use:   "describe-rule",
	Short: "Describes a rule for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeRuleCmd).Standalone()

	connect_describeRuleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeRuleCmd.Flags().String("rule-id", "", "A unique identifier for the rule.")
	connect_describeRuleCmd.MarkFlagRequired("instance-id")
	connect_describeRuleCmd.MarkFlagRequired("rule-id")
	connectCmd.AddCommand(connect_describeRuleCmd)
}
