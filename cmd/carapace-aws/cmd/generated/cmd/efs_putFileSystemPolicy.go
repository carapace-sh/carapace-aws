package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_putFileSystemPolicyCmd = &cobra.Command{
	Use:   "put-file-system-policy",
	Short: "Applies an Amazon EFS `FileSystemPolicy` to an Amazon EFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_putFileSystemPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_putFileSystemPolicyCmd).Standalone()

		efs_putFileSystemPolicyCmd.Flags().String("bypass-policy-lockout-safety-check", "", "(Optional) A boolean that specifies whether or not to bypass the `FileSystemPolicy` lockout safety check.")
		efs_putFileSystemPolicyCmd.Flags().String("file-system-id", "", "The ID of the EFS file system that you want to create or update the `FileSystemPolicy` for.")
		efs_putFileSystemPolicyCmd.Flags().String("policy", "", "The `FileSystemPolicy` that you're creating.")
		efs_putFileSystemPolicyCmd.MarkFlagRequired("file-system-id")
		efs_putFileSystemPolicyCmd.MarkFlagRequired("policy")
	})
	efsCmd.AddCommand(efs_putFileSystemPolicyCmd)
}
