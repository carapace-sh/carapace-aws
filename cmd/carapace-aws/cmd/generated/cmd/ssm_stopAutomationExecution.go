package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_stopAutomationExecutionCmd = &cobra.Command{
	Use:   "stop-automation-execution",
	Short: "Stop an Automation that is currently running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_stopAutomationExecutionCmd).Standalone()

	ssm_stopAutomationExecutionCmd.Flags().String("automation-execution-id", "", "The execution ID of the Automation to stop.")
	ssm_stopAutomationExecutionCmd.Flags().String("type", "", "The stop request type.")
	ssm_stopAutomationExecutionCmd.MarkFlagRequired("automation-execution-id")
	ssmCmd.AddCommand(ssm_stopAutomationExecutionCmd)
}
