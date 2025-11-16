package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_modifyMountTargetSecurityGroupsCmd = &cobra.Command{
	Use:   "modify-mount-target-security-groups",
	Short: "Modifies the set of security groups in effect for a mount target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_modifyMountTargetSecurityGroupsCmd).Standalone()

	efs_modifyMountTargetSecurityGroupsCmd.Flags().String("mount-target-id", "", "The ID of the mount target whose security groups you want to modify.")
	efs_modifyMountTargetSecurityGroupsCmd.Flags().String("security-groups", "", "An array of VPC security group IDs.")
	efs_modifyMountTargetSecurityGroupsCmd.MarkFlagRequired("mount-target-id")
	efsCmd.AddCommand(efs_modifyMountTargetSecurityGroupsCmd)
}
