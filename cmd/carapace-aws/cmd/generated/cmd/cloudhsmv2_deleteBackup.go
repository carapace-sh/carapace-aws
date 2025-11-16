package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_deleteBackupCmd = &cobra.Command{
	Use:   "delete-backup",
	Short: "Deletes a specified CloudHSM backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_deleteBackupCmd).Standalone()

	cloudhsmv2_deleteBackupCmd.Flags().String("backup-id", "", "The ID of the backup to be deleted.")
	cloudhsmv2_deleteBackupCmd.MarkFlagRequired("backup-id")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_deleteBackupCmd)
}
