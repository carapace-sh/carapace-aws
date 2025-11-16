package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateRecoveryPointIndexSettingsCmd = &cobra.Command{
	Use:   "update-recovery-point-index-settings",
	Short: "This operation updates the settings of a recovery point index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateRecoveryPointIndexSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_updateRecoveryPointIndexSettingsCmd).Standalone()

		backup_updateRecoveryPointIndexSettingsCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_updateRecoveryPointIndexSettingsCmd.Flags().String("iam-role-arn", "", "This specifies the IAM role ARN used for this operation.")
		backup_updateRecoveryPointIndexSettingsCmd.Flags().String("index", "", "Index can have 1 of 2 possible values, either `ENABLED` or `DISABLED`.")
		backup_updateRecoveryPointIndexSettingsCmd.Flags().String("recovery-point-arn", "", "An ARN that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
		backup_updateRecoveryPointIndexSettingsCmd.MarkFlagRequired("backup-vault-name")
		backup_updateRecoveryPointIndexSettingsCmd.MarkFlagRequired("index")
		backup_updateRecoveryPointIndexSettingsCmd.MarkFlagRequired("recovery-point-arn")
	})
	backupCmd.AddCommand(backup_updateRecoveryPointIndexSettingsCmd)
}
