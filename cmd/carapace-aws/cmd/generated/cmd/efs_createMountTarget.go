package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_createMountTargetCmd = &cobra.Command{
	Use:   "create-mount-target",
	Short: "Creates a mount target for a file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_createMountTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_createMountTargetCmd).Standalone()

		efs_createMountTargetCmd.Flags().String("file-system-id", "", "The ID of the file system for which to create the mount target.")
		efs_createMountTargetCmd.Flags().String("ip-address", "", "If the IP address type for the mount target is IPv4, then specify the IPv4 address within the address range of the specified subnet.")
		efs_createMountTargetCmd.Flags().String("ip-address-type", "", "Specify the type of IP address of the mount target you are creating.")
		efs_createMountTargetCmd.Flags().String("ipv6-address", "", "If the IP address type for the mount target is IPv6, then specify the IPv6 address within the address range of the specified subnet.")
		efs_createMountTargetCmd.Flags().String("security-groups", "", "VPC security group IDs, of the form `sg-xxxxxxxx`.")
		efs_createMountTargetCmd.Flags().String("subnet-id", "", "The ID of the subnet to add the mount target in.")
		efs_createMountTargetCmd.MarkFlagRequired("file-system-id")
		efs_createMountTargetCmd.MarkFlagRequired("subnet-id")
	})
	efsCmd.AddCommand(efs_createMountTargetCmd)
}
