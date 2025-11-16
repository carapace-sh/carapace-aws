package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeRestoreJobCmd = &cobra.Command{
	Use:   "describe-restore-job",
	Short: "Returns metadata associated with a restore job that is specified by a job ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeRestoreJobCmd).Standalone()

	backup_describeRestoreJobCmd.Flags().String("restore-job-id", "", "Uniquely identifies the job that restores a recovery point.")
	backup_describeRestoreJobCmd.MarkFlagRequired("restore-job-id")
	backupCmd.AddCommand(backup_describeRestoreJobCmd)
}
