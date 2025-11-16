package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_provisionPublicIpv4PoolCidrCmd = &cobra.Command{
	Use:   "provision-public-ipv4-pool-cidr",
	Short: "Provision a CIDR to a public IPv4 pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_provisionPublicIpv4PoolCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_provisionPublicIpv4PoolCidrCmd).Standalone()

		ec2_provisionPublicIpv4PoolCidrCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_provisionPublicIpv4PoolCidrCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool you would like to use to allocate this CIDR.")
		ec2_provisionPublicIpv4PoolCidrCmd.Flags().String("netmask-length", "", "The netmask length of the CIDR you would like to allocate to the public IPv4 pool.")
		ec2_provisionPublicIpv4PoolCidrCmd.Flags().String("network-border-group", "", "The Availability Zone (AZ) or Local Zone (LZ) network border group that the resource that the IP address is assigned to is in.")
		ec2_provisionPublicIpv4PoolCidrCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_provisionPublicIpv4PoolCidrCmd.Flags().String("pool-id", "", "The ID of the public IPv4 pool you would like to use for this CIDR.")
		ec2_provisionPublicIpv4PoolCidrCmd.MarkFlagRequired("ipam-pool-id")
		ec2_provisionPublicIpv4PoolCidrCmd.MarkFlagRequired("netmask-length")
		ec2_provisionPublicIpv4PoolCidrCmd.Flag("no-dry-run").Hidden = true
		ec2_provisionPublicIpv4PoolCidrCmd.MarkFlagRequired("pool-id")
	})
	ec2Cmd.AddCommand(ec2_provisionPublicIpv4PoolCidrCmd)
}
