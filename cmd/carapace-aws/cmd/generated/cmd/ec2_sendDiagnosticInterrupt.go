package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_sendDiagnosticInterruptCmd = &cobra.Command{
	Use:   "send-diagnostic-interrupt",
	Short: "Sends a diagnostic interrupt to the specified Amazon EC2 instance to trigger a *kernel panic* (on Linux instances), or a *blue screen*/*stop error* (on Windows instances).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_sendDiagnosticInterruptCmd).Standalone()

	ec2_sendDiagnosticInterruptCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_sendDiagnosticInterruptCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_sendDiagnosticInterruptCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_sendDiagnosticInterruptCmd.MarkFlagRequired("instance-id")
	ec2_sendDiagnosticInterruptCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_sendDiagnosticInterruptCmd)
}
