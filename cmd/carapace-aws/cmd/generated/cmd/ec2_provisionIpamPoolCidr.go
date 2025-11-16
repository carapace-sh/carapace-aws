package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_provisionIpamPoolCidrCmd = &cobra.Command{
	Use:   "provision-ipam-pool-cidr",
	Short: "Provision a CIDR to an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_provisionIpamPoolCidrCmd).Standalone()

	ec2_provisionIpamPoolCidrCmd.Flags().String("cidr", "", "The CIDR you want to assign to the IPAM pool.")
	ec2_provisionIpamPoolCidrCmd.Flags().String("cidr-authorization-context", "", "A signed document that proves that you are authorized to bring a specified IP address range to Amazon using BYOIP.")
	ec2_provisionIpamPoolCidrCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_provisionIpamPoolCidrCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_provisionIpamPoolCidrCmd.Flags().String("ipam-external-resource-verification-token-id", "", "Verification token ID.")
	ec2_provisionIpamPoolCidrCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool to which you want to assign a CIDR.")
	ec2_provisionIpamPoolCidrCmd.Flags().String("netmask-length", "", "The netmask length of the CIDR you'd like to provision to a pool.")
	ec2_provisionIpamPoolCidrCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_provisionIpamPoolCidrCmd.Flags().String("verification-method", "", "The method for verifying control of a public IP address range.")
	ec2_provisionIpamPoolCidrCmd.MarkFlagRequired("ipam-pool-id")
	ec2_provisionIpamPoolCidrCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_provisionIpamPoolCidrCmd)
}
