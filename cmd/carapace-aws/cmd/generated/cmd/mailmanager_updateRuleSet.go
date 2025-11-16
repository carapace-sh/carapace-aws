package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_updateRuleSetCmd = &cobra.Command{
	Use:   "update-rule-set",
	Short: "Update attributes of an already provisioned rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_updateRuleSetCmd).Standalone()

	mailmanager_updateRuleSetCmd.Flags().String("rule-set-id", "", "The identifier of a rule set you want to update.")
	mailmanager_updateRuleSetCmd.Flags().String("rule-set-name", "", "A user-friendly name for the rule set resource.")
	mailmanager_updateRuleSetCmd.Flags().String("rules", "", "A new set of rules to replace the current rules of the rule set—these rules will override all the rules of the rule set.")
	mailmanager_updateRuleSetCmd.MarkFlagRequired("rule-set-id")
	mailmanagerCmd.AddCommand(mailmanager_updateRuleSetCmd)
}
