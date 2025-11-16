package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateRecoveryPointLifecycleCmd = &cobra.Command{
	Use:   "update-recovery-point-lifecycle",
	Short: "Sets the transition lifecycle of a recovery point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateRecoveryPointLifecycleCmd).Standalone()

	backup_updateRecoveryPointLifecycleCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_updateRecoveryPointLifecycleCmd.Flags().String("lifecycle", "", "The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.")
	backup_updateRecoveryPointLifecycleCmd.Flags().String("recovery-point-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
	backup_updateRecoveryPointLifecycleCmd.MarkFlagRequired("backup-vault-name")
	backup_updateRecoveryPointLifecycleCmd.MarkFlagRequired("recovery-point-arn")
	backupCmd.AddCommand(backup_updateRecoveryPointLifecycleCmd)
}
