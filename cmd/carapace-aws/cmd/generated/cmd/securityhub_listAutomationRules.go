package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listAutomationRulesCmd = &cobra.Command{
	Use:   "list-automation-rules",
	Short: "A list of automation rules and their metadata for the calling account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listAutomationRulesCmd).Standalone()

	securityhub_listAutomationRulesCmd.Flags().String("max-results", "", "The maximum number of rules to return in the response.")
	securityhub_listAutomationRulesCmd.Flags().String("next-token", "", "A token to specify where to start paginating the response.")
	securityhubCmd.AddCommand(securityhub_listAutomationRulesCmd)
}
