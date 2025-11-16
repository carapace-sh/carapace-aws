package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createBackupCmd = &cobra.Command{
	Use:   "create-backup",
	Short: "Creates a backup of an existing Amazon FSx for Windows File Server file system, Amazon FSx for Lustre file system, Amazon FSx for NetApp ONTAP volume, or Amazon FSx for OpenZFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createBackupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_createBackupCmd).Standalone()

		fsx_createBackupCmd.Flags().String("client-request-token", "", "(Optional) A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent creation.")
		fsx_createBackupCmd.Flags().String("file-system-id", "", "The ID of the file system to back up.")
		fsx_createBackupCmd.Flags().String("tags", "", "(Optional) The tags to apply to the backup at backup creation.")
		fsx_createBackupCmd.Flags().String("volume-id", "", "(Optional) The ID of the FSx for ONTAP volume to back up.")
	})
	fsxCmd.AddCommand(fsx_createBackupCmd)
}
