package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_startAutomationExecutionCmd = &cobra.Command{
	Use:   "start-automation-execution",
	Short: "Initiates execution of an Automation runbook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_startAutomationExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_startAutomationExecutionCmd).Standalone()

		ssm_startAutomationExecutionCmd.Flags().String("alarm-configuration", "", "The CloudWatch alarm you want to apply to your automation.")
		ssm_startAutomationExecutionCmd.Flags().String("client-token", "", "User-provided idempotency token.")
		ssm_startAutomationExecutionCmd.Flags().String("document-name", "", "The name of the SSM document to run.")
		ssm_startAutomationExecutionCmd.Flags().String("document-version", "", "The version of the Automation runbook to use for this execution.")
		ssm_startAutomationExecutionCmd.Flags().String("max-concurrency", "", "The maximum number of targets allowed to run this task in parallel.")
		ssm_startAutomationExecutionCmd.Flags().String("max-errors", "", "The number of errors that are allowed before the system stops running the automation on additional targets.")
		ssm_startAutomationExecutionCmd.Flags().String("mode", "", "The execution mode of the automation.")
		ssm_startAutomationExecutionCmd.Flags().String("parameters", "", "A key-value map of execution parameters, which match the declared parameters in the Automation runbook.")
		ssm_startAutomationExecutionCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
		ssm_startAutomationExecutionCmd.Flags().String("target-locations", "", "A location is a combination of Amazon Web Services Regions and/or Amazon Web Services accounts where you want to run the automation.")
		ssm_startAutomationExecutionCmd.Flags().String("target-locations-url", "", "Specify a publicly accessible URL for a file that contains the `TargetLocations` body.")
		ssm_startAutomationExecutionCmd.Flags().String("target-maps", "", "A key-value mapping of document parameters to target resources.")
		ssm_startAutomationExecutionCmd.Flags().String("target-parameter-name", "", "The name of the parameter used as the target resource for the rate-controlled execution.")
		ssm_startAutomationExecutionCmd.Flags().String("targets", "", "A key-value mapping to target resources.")
		ssm_startAutomationExecutionCmd.MarkFlagRequired("document-name")
	})
	ssmCmd.AddCommand(ssm_startAutomationExecutionCmd)
}
