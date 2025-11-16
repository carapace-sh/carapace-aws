package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getAutomationExecutionCmd = &cobra.Command{
	Use:   "get-automation-execution",
	Short: "Get detailed information about a particular Automation execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getAutomationExecutionCmd).Standalone()

	ssm_getAutomationExecutionCmd.Flags().String("automation-execution-id", "", "The unique identifier for an existing automation execution to examine.")
	ssm_getAutomationExecutionCmd.MarkFlagRequired("automation-execution-id")
	ssmCmd.AddCommand(ssm_getAutomationExecutionCmd)
}
