package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createRuleSetCmd = &cobra.Command{
	Use:   "create-rule-set",
	Short: "Provision a new rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createRuleSetCmd).Standalone()

	mailmanager_createRuleSetCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
	mailmanager_createRuleSetCmd.Flags().String("rule-set-name", "", "A user-friendly name for the rule set.")
	mailmanager_createRuleSetCmd.Flags().String("rules", "", "Conditional rules that are evaluated for determining actions on email.")
	mailmanager_createRuleSetCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
	mailmanager_createRuleSetCmd.MarkFlagRequired("rule-set-name")
	mailmanager_createRuleSetCmd.MarkFlagRequired("rules")
	mailmanagerCmd.AddCommand(mailmanager_createRuleSetCmd)
}
