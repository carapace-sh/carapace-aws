package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelImportTaskCmd = &cobra.Command{
	Use:   "cancel-import-task",
	Short: "Cancels an in-process import virtual machine or import snapshot task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelImportTaskCmd).Standalone()

	ec2_cancelImportTaskCmd.Flags().String("cancel-reason", "", "The reason for canceling the task.")
	ec2_cancelImportTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelImportTaskCmd.Flags().String("import-task-id", "", "The ID of the import image or import snapshot task to be canceled.")
	ec2_cancelImportTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelImportTaskCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_cancelImportTaskCmd)
}
