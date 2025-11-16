package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_putRestoreValidationResultCmd = &cobra.Command{
	Use:   "put-restore-validation-result",
	Short: "This request allows you to send your independent self-run restore test validation results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_putRestoreValidationResultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_putRestoreValidationResultCmd).Standalone()

		backup_putRestoreValidationResultCmd.Flags().String("restore-job-id", "", "This is a unique identifier of a restore job within Backup.")
		backup_putRestoreValidationResultCmd.Flags().String("validation-status", "", "The status of your restore validation.")
		backup_putRestoreValidationResultCmd.Flags().String("validation-status-message", "", "This is an optional message string you can input to describe the validation status for the restore test validation.")
		backup_putRestoreValidationResultCmd.MarkFlagRequired("restore-job-id")
		backup_putRestoreValidationResultCmd.MarkFlagRequired("validation-status")
	})
	backupCmd.AddCommand(backup_putRestoreValidationResultCmd)
}
