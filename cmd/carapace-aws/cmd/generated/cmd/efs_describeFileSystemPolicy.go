package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeFileSystemPolicyCmd = &cobra.Command{
	Use:   "describe-file-system-policy",
	Short: "Returns the `FileSystemPolicy` for the specified EFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeFileSystemPolicyCmd).Standalone()

	efs_describeFileSystemPolicyCmd.Flags().String("file-system-id", "", "Specifies which EFS file system to retrieve the `FileSystemPolicy` for.")
	efs_describeFileSystemPolicyCmd.MarkFlagRequired("file-system-id")
	efsCmd.AddCommand(efs_describeFileSystemPolicyCmd)
}
