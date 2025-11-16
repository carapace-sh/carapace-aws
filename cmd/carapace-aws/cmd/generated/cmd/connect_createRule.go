package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "Creates a rule for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createRuleCmd).Standalone()

		connect_createRuleCmd.Flags().String("actions", "", "A list of actions to be run when the rule is triggered.")
		connect_createRuleCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_createRuleCmd.Flags().String("function", "", "The conditions of the rule.")
		connect_createRuleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createRuleCmd.Flags().String("name", "", "A unique name for the rule.")
		connect_createRuleCmd.Flags().String("publish-status", "", "The publish status of the rule.")
		connect_createRuleCmd.Flags().String("trigger-event-source", "", "The event source to trigger the rule.")
		connect_createRuleCmd.MarkFlagRequired("actions")
		connect_createRuleCmd.MarkFlagRequired("function")
		connect_createRuleCmd.MarkFlagRequired("instance-id")
		connect_createRuleCmd.MarkFlagRequired("name")
		connect_createRuleCmd.MarkFlagRequired("publish-status")
		connect_createRuleCmd.MarkFlagRequired("trigger-event-source")
	})
	connectCmd.AddCommand(connect_createRuleCmd)
}
