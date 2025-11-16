package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelConversionTaskCmd = &cobra.Command{
	Use:   "cancel-conversion-task",
	Short: "Cancels an active conversion task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelConversionTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_cancelConversionTaskCmd).Standalone()

		ec2_cancelConversionTaskCmd.Flags().String("conversion-task-id", "", "The ID of the conversion task.")
		ec2_cancelConversionTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelConversionTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelConversionTaskCmd.Flags().String("reason-message", "", "The reason for canceling the conversion task.")
		ec2_cancelConversionTaskCmd.MarkFlagRequired("conversion-task-id")
		ec2_cancelConversionTaskCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_cancelConversionTaskCmd)
}
