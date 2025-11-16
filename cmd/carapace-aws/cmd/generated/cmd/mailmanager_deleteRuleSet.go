package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteRuleSetCmd = &cobra.Command{
	Use:   "delete-rule-set",
	Short: "Delete a rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_deleteRuleSetCmd).Standalone()

		mailmanager_deleteRuleSetCmd.Flags().String("rule-set-id", "", "The identifier of an existing rule set resource to delete.")
		mailmanager_deleteRuleSetCmd.MarkFlagRequired("rule-set-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_deleteRuleSetCmd)
}
