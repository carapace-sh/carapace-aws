package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listAutomationRulesV2Cmd = &cobra.Command{
	Use:   "list-automation-rules-v2",
	Short: "Returns a list of automation rules and metadata for the calling account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listAutomationRulesV2Cmd).Standalone()

	securityhub_listAutomationRulesV2Cmd.Flags().String("max-results", "", "The maximum number of results to return.")
	securityhub_listAutomationRulesV2Cmd.Flags().String("next-token", "", "The token required for pagination.")
	securityhubCmd.AddCommand(securityhub_listAutomationRulesV2Cmd)
}
