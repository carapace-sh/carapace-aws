package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_putBackupPolicyCmd = &cobra.Command{
	Use:   "put-backup-policy",
	Short: "Updates the file system's backup policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_putBackupPolicyCmd).Standalone()

	efs_putBackupPolicyCmd.Flags().String("backup-policy", "", "The backup policy included in the `PutBackupPolicy` request.")
	efs_putBackupPolicyCmd.Flags().String("file-system-id", "", "Specifies which EFS file system to update the backup policy for.")
	efs_putBackupPolicyCmd.MarkFlagRequired("backup-policy")
	efs_putBackupPolicyCmd.MarkFlagRequired("file-system-id")
	efsCmd.AddCommand(efs_putBackupPolicyCmd)
}
