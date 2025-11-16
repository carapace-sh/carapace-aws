package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_advertiseByoipCidrCmd = &cobra.Command{
	Use:   "advertise-byoip-cidr",
	Short: "Advertises an IPv4 or IPv6 address range that is provisioned for use with your Amazon Web Services resources through bring your own IP addresses (BYOIP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_advertiseByoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_advertiseByoipCidrCmd).Standalone()

		ec2_advertiseByoipCidrCmd.Flags().String("asn", "", "The public 2-byte or 4-byte ASN that you want to advertise.")
		ec2_advertiseByoipCidrCmd.Flags().String("cidr", "", "The address range, in CIDR notation.")
		ec2_advertiseByoipCidrCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_advertiseByoipCidrCmd.Flags().String("network-border-group", "", "If you have [Local Zones](https://docs.aws.amazon.com/local-zones/latest/ug/how-local-zones-work.html) enabled, you can choose a network border group for Local Zones when you provision and advertise a BYOIPv4 CIDR.")
		ec2_advertiseByoipCidrCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_advertiseByoipCidrCmd.MarkFlagRequired("cidr")
		ec2_advertiseByoipCidrCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_advertiseByoipCidrCmd)
}
