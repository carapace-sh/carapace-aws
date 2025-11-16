package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_allocateIpamPoolCidrCmd = &cobra.Command{
	Use:   "allocate-ipam-pool-cidr",
	Short: "Allocate a CIDR from an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_allocateIpamPoolCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_allocateIpamPoolCidrCmd).Standalone()

		ec2_allocateIpamPoolCidrCmd.Flags().String("allowed-cidrs", "", "Include a particular CIDR range that can be returned by the pool.")
		ec2_allocateIpamPoolCidrCmd.Flags().String("cidr", "", "The CIDR you would like to allocate from the IPAM pool.")
		ec2_allocateIpamPoolCidrCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_allocateIpamPoolCidrCmd.Flags().String("description", "", "A description for the allocation.")
		ec2_allocateIpamPoolCidrCmd.Flags().String("disallowed-cidrs", "", "Exclude a particular CIDR range from being returned by the pool.")
		ec2_allocateIpamPoolCidrCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_allocateIpamPoolCidrCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool from which you would like to allocate a CIDR.")
		ec2_allocateIpamPoolCidrCmd.Flags().String("netmask-length", "", "The netmask length of the CIDR you would like to allocate from the IPAM pool.")
		ec2_allocateIpamPoolCidrCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_allocateIpamPoolCidrCmd.Flags().Bool("no-preview-next-cidr", false, "A preview of the next available CIDR in a pool.")
		ec2_allocateIpamPoolCidrCmd.Flags().Bool("preview-next-cidr", false, "A preview of the next available CIDR in a pool.")
		ec2_allocateIpamPoolCidrCmd.MarkFlagRequired("ipam-pool-id")
		ec2_allocateIpamPoolCidrCmd.Flag("no-dry-run").Hidden = true
		ec2_allocateIpamPoolCidrCmd.Flag("no-preview-next-cidr").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_allocateIpamPoolCidrCmd)
}
