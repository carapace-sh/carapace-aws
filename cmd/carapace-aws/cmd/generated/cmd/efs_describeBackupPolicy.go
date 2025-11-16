package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeBackupPolicyCmd = &cobra.Command{
	Use:   "describe-backup-policy",
	Short: "Returns the backup policy for the specified EFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeBackupPolicyCmd).Standalone()

	efs_describeBackupPolicyCmd.Flags().String("file-system-id", "", "Specifies which EFS file system for which to retrieve the `BackupPolicy`.")
	efs_describeBackupPolicyCmd.MarkFlagRequired("file-system-id")
	efsCmd.AddCommand(efs_describeBackupPolicyCmd)
}
