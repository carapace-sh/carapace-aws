package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeMountTargetSecurityGroupsCmd = &cobra.Command{
	Use:   "describe-mount-target-security-groups",
	Short: "Returns the security groups currently in effect for a mount target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeMountTargetSecurityGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_describeMountTargetSecurityGroupsCmd).Standalone()

		efs_describeMountTargetSecurityGroupsCmd.Flags().String("mount-target-id", "", "The ID of the mount target whose security groups you want to retrieve.")
		efs_describeMountTargetSecurityGroupsCmd.MarkFlagRequired("mount-target-id")
	})
	efsCmd.AddCommand(efs_describeMountTargetSecurityGroupsCmd)
}
