package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getRuleSetCmd = &cobra.Command{
	Use:   "get-rule-set",
	Short: "Fetch attributes of a rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getRuleSetCmd).Standalone()

	mailmanager_getRuleSetCmd.Flags().String("rule-set-id", "", "The identifier of an existing rule set to be retrieved.")
	mailmanager_getRuleSetCmd.MarkFlagRequired("rule-set-id")
	mailmanagerCmd.AddCommand(mailmanager_getRuleSetCmd)
}
