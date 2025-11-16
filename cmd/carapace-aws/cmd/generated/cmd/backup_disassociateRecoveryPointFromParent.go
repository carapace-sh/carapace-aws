package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_disassociateRecoveryPointFromParentCmd = &cobra.Command{
	Use:   "disassociate-recovery-point-from-parent",
	Short: "This action to a specific child (nested) recovery point removes the relationship between the specified recovery point and its parent (composite) recovery point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_disassociateRecoveryPointFromParentCmd).Standalone()

	backup_disassociateRecoveryPointFromParentCmd.Flags().String("backup-vault-name", "", "The name of a logical container where the child (nested) recovery point is stored.")
	backup_disassociateRecoveryPointFromParentCmd.Flags().String("recovery-point-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the child (nested) recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.`")
	backup_disassociateRecoveryPointFromParentCmd.MarkFlagRequired("backup-vault-name")
	backup_disassociateRecoveryPointFromParentCmd.MarkFlagRequired("recovery-point-arn")
	backupCmd.AddCommand(backup_disassociateRecoveryPointFromParentCmd)
}
