package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_deleteMountTargetCmd = &cobra.Command{
	Use:   "delete-mount-target",
	Short: "Deletes the specified mount target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_deleteMountTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_deleteMountTargetCmd).Standalone()

		efs_deleteMountTargetCmd.Flags().String("mount-target-id", "", "The ID of the mount target to delete (String).")
		efs_deleteMountTargetCmd.MarkFlagRequired("mount-target-id")
	})
	efsCmd.AddCommand(efs_deleteMountTargetCmd)
}
