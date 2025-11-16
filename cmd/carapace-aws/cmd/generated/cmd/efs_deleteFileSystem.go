package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_deleteFileSystemCmd = &cobra.Command{
	Use:   "delete-file-system",
	Short: "Deletes a file system, permanently severing access to its contents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_deleteFileSystemCmd).Standalone()

	efs_deleteFileSystemCmd.Flags().String("file-system-id", "", "The ID of the file system you want to delete.")
	efs_deleteFileSystemCmd.MarkFlagRequired("file-system-id")
	efsCmd.AddCommand(efs_deleteFileSystemCmd)
}
