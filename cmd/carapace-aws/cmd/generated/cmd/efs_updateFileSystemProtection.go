package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_updateFileSystemProtectionCmd = &cobra.Command{
	Use:   "update-file-system-protection",
	Short: "Updates protection on the file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_updateFileSystemProtectionCmd).Standalone()

	efs_updateFileSystemProtectionCmd.Flags().String("file-system-id", "", "The ID of the file system to update.")
	efs_updateFileSystemProtectionCmd.Flags().String("replication-overwrite-protection", "", "The status of the file system's replication overwrite protection.")
	efs_updateFileSystemProtectionCmd.MarkFlagRequired("file-system-id")
	efsCmd.AddCommand(efs_updateFileSystemProtectionCmd)
}
