package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_moveByoipCidrToIpamCmd = &cobra.Command{
	Use:   "move-byoip-cidr-to-ipam",
	Short: "Move a BYOIPv4 CIDR to IPAM from a public IPv4 pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_moveByoipCidrToIpamCmd).Standalone()

	ec2_moveByoipCidrToIpamCmd.Flags().String("cidr", "", "The BYOIP CIDR.")
	ec2_moveByoipCidrToIpamCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_moveByoipCidrToIpamCmd.Flags().String("ipam-pool-id", "", "The IPAM pool ID.")
	ec2_moveByoipCidrToIpamCmd.Flags().String("ipam-pool-owner", "", "The Amazon Web Services account ID of the owner of the IPAM pool.")
	ec2_moveByoipCidrToIpamCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_moveByoipCidrToIpamCmd.MarkFlagRequired("cidr")
	ec2_moveByoipCidrToIpamCmd.MarkFlagRequired("ipam-pool-id")
	ec2_moveByoipCidrToIpamCmd.MarkFlagRequired("ipam-pool-owner")
	ec2_moveByoipCidrToIpamCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_moveByoipCidrToIpamCmd)
}
