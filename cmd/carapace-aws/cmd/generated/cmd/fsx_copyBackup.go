package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_copyBackupCmd = &cobra.Command{
	Use:   "copy-backup",
	Short: "Copies an existing backup within the same Amazon Web Services account to another Amazon Web Services Region (cross-Region copy) or within the same Amazon Web Services Region (in-Region copy).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_copyBackupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_copyBackupCmd).Standalone()

		fsx_copyBackupCmd.Flags().String("client-request-token", "", "")
		fsx_copyBackupCmd.Flags().String("copy-tags", "", "A Boolean flag indicating whether tags from the source backup should be copied to the backup copy.")
		fsx_copyBackupCmd.Flags().String("kms-key-id", "", "")
		fsx_copyBackupCmd.Flags().String("source-backup-id", "", "The ID of the source backup.")
		fsx_copyBackupCmd.Flags().String("source-region", "", "The source Amazon Web Services Region of the backup.")
		fsx_copyBackupCmd.Flags().String("tags", "", "")
		fsx_copyBackupCmd.MarkFlagRequired("source-backup-id")
	})
	fsxCmd.AddCommand(fsx_copyBackupCmd)
}
