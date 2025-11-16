package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_sendAutomationSignalCmd = &cobra.Command{
	Use:   "send-automation-signal",
	Short: "Sends a signal to an Automation execution to change the current behavior or status of the execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_sendAutomationSignalCmd).Standalone()

	ssm_sendAutomationSignalCmd.Flags().String("automation-execution-id", "", "The unique identifier for an existing Automation execution that you want to send the signal to.")
	ssm_sendAutomationSignalCmd.Flags().String("payload", "", "The data sent with the signal.")
	ssm_sendAutomationSignalCmd.Flags().String("signal-type", "", "The type of signal to send to an Automation execution.")
	ssm_sendAutomationSignalCmd.MarkFlagRequired("automation-execution-id")
	ssm_sendAutomationSignalCmd.MarkFlagRequired("signal-type")
	ssmCmd.AddCommand(ssm_sendAutomationSignalCmd)
}
