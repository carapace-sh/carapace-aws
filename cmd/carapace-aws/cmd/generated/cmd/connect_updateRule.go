package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateRuleCmd = &cobra.Command{
	Use:   "update-rule",
	Short: "Updates a rule for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateRuleCmd).Standalone()

		connect_updateRuleCmd.Flags().String("actions", "", "A list of actions to be run when the rule is triggered.")
		connect_updateRuleCmd.Flags().String("function", "", "The conditions of the rule.")
		connect_updateRuleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateRuleCmd.Flags().String("name", "", "The name of the rule.")
		connect_updateRuleCmd.Flags().String("publish-status", "", "The publish status of the rule.")
		connect_updateRuleCmd.Flags().String("rule-id", "", "A unique identifier for the rule.")
		connect_updateRuleCmd.MarkFlagRequired("actions")
		connect_updateRuleCmd.MarkFlagRequired("function")
		connect_updateRuleCmd.MarkFlagRequired("instance-id")
		connect_updateRuleCmd.MarkFlagRequired("name")
		connect_updateRuleCmd.MarkFlagRequired("publish-status")
		connect_updateRuleCmd.MarkFlagRequired("rule-id")
	})
	connectCmd.AddCommand(connect_updateRuleCmd)
}
