package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Deletes a rule for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteRuleCmd).Standalone()

		connect_deleteRuleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteRuleCmd.Flags().String("rule-id", "", "A unique identifier for the rule.")
		connect_deleteRuleCmd.MarkFlagRequired("instance-id")
		connect_deleteRuleCmd.MarkFlagRequired("rule-id")
	})
	connectCmd.AddCommand(connect_deleteRuleCmd)
}
