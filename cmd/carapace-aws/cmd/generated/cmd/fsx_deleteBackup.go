package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_deleteBackupCmd = &cobra.Command{
	Use:   "delete-backup",
	Short: "Deletes an Amazon FSx backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_deleteBackupCmd).Standalone()

	fsx_deleteBackupCmd.Flags().String("backup-id", "", "The ID of the backup that you want to delete.")
	fsx_deleteBackupCmd.Flags().String("client-request-token", "", "A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent deletion.")
	fsx_deleteBackupCmd.MarkFlagRequired("backup-id")
	fsxCmd.AddCommand(fsx_deleteBackupCmd)
}
