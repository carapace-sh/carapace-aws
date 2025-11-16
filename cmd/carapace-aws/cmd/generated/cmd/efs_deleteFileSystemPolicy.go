package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_deleteFileSystemPolicyCmd = &cobra.Command{
	Use:   "delete-file-system-policy",
	Short: "Deletes the `FileSystemPolicy` for the specified file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_deleteFileSystemPolicyCmd).Standalone()

	efs_deleteFileSystemPolicyCmd.Flags().String("file-system-id", "", "Specifies the EFS file system for which to delete the `FileSystemPolicy`.")
	efs_deleteFileSystemPolicyCmd.MarkFlagRequired("file-system-id")
	efsCmd.AddCommand(efs_deleteFileSystemPolicyCmd)
}
