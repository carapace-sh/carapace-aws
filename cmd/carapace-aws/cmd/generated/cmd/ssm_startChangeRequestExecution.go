package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_startChangeRequestExecutionCmd = &cobra.Command{
	Use:   "start-change-request-execution",
	Short: "Amazon Web Services Systems Manager Change Manager will no longer be open to new customers starting November 7, 2025. If you would like to use Change Manager, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_startChangeRequestExecutionCmd).Standalone()

	ssm_startChangeRequestExecutionCmd.Flags().Bool("auto-approve", false, "Indicates whether the change request can be approved automatically without the need for manual approvals.")
	ssm_startChangeRequestExecutionCmd.Flags().String("change-details", "", "User-provided details about the change.")
	ssm_startChangeRequestExecutionCmd.Flags().String("change-request-name", "", "The name of the change request associated with the runbook workflow to be run.")
	ssm_startChangeRequestExecutionCmd.Flags().String("client-token", "", "The user-provided idempotency token.")
	ssm_startChangeRequestExecutionCmd.Flags().String("document-name", "", "The name of the change template document to run during the runbook workflow.")
	ssm_startChangeRequestExecutionCmd.Flags().String("document-version", "", "The version of the change template document to run during the runbook workflow.")
	ssm_startChangeRequestExecutionCmd.Flags().Bool("no-auto-approve", false, "Indicates whether the change request can be approved automatically without the need for manual approvals.")
	ssm_startChangeRequestExecutionCmd.Flags().String("parameters", "", "A key-value map of parameters that match the declared parameters in the change template document.")
	ssm_startChangeRequestExecutionCmd.Flags().String("runbooks", "", "Information about the Automation runbooks that are run during the runbook workflow.")
	ssm_startChangeRequestExecutionCmd.Flags().String("scheduled-end-time", "", "The time that the requester expects the runbook workflow related to the change request to complete.")
	ssm_startChangeRequestExecutionCmd.Flags().String("scheduled-time", "", "The date and time specified in the change request to run the Automation runbooks.")
	ssm_startChangeRequestExecutionCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
	ssm_startChangeRequestExecutionCmd.MarkFlagRequired("document-name")
	ssm_startChangeRequestExecutionCmd.Flag("no-auto-approve").Hidden = true
	ssm_startChangeRequestExecutionCmd.MarkFlagRequired("runbooks")
	ssmCmd.AddCommand(ssm_startChangeRequestExecutionCmd)
}
