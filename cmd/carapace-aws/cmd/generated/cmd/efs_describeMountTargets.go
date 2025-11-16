package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeMountTargetsCmd = &cobra.Command{
	Use:   "describe-mount-targets",
	Short: "Returns the descriptions of all the current mount targets, or a specific mount target, for a file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeMountTargetsCmd).Standalone()

	efs_describeMountTargetsCmd.Flags().String("access-point-id", "", "(Optional) The ID of the access point whose mount targets that you want to list.")
	efs_describeMountTargetsCmd.Flags().String("file-system-id", "", "(Optional) ID of the file system whose mount targets you want to list (String).")
	efs_describeMountTargetsCmd.Flags().String("marker", "", "(Optional) Opaque pagination token returned from a previous `DescribeMountTargets` operation (String).")
	efs_describeMountTargetsCmd.Flags().String("max-items", "", "(Optional) Maximum number of mount targets to return in the response.")
	efs_describeMountTargetsCmd.Flags().String("mount-target-id", "", "(Optional) ID of the mount target that you want to have described (String).")
	efsCmd.AddCommand(efs_describeMountTargetsCmd)
}
