package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_releaseIpamPoolAllocationCmd = &cobra.Command{
	Use:   "release-ipam-pool-allocation",
	Short: "Release an allocation within an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_releaseIpamPoolAllocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_releaseIpamPoolAllocationCmd).Standalone()

		ec2_releaseIpamPoolAllocationCmd.Flags().String("cidr", "", "The CIDR of the allocation you want to release.")
		ec2_releaseIpamPoolAllocationCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_releaseIpamPoolAllocationCmd.Flags().String("ipam-pool-allocation-id", "", "The ID of the allocation.")
		ec2_releaseIpamPoolAllocationCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool which contains the allocation you want to release.")
		ec2_releaseIpamPoolAllocationCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_releaseIpamPoolAllocationCmd.MarkFlagRequired("cidr")
		ec2_releaseIpamPoolAllocationCmd.MarkFlagRequired("ipam-pool-allocation-id")
		ec2_releaseIpamPoolAllocationCmd.MarkFlagRequired("ipam-pool-id")
		ec2_releaseIpamPoolAllocationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_releaseIpamPoolAllocationCmd)
}
