package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_provisionByoipCidrCmd = &cobra.Command{
	Use:   "provision-byoip-cidr",
	Short: "Provisions an IPv4 or IPv6 address range for use with your Amazon Web Services resources through bring your own IP addresses (BYOIP) and creates a corresponding address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_provisionByoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_provisionByoipCidrCmd).Standalone()

		ec2_provisionByoipCidrCmd.Flags().String("cidr", "", "The public IPv4 or IPv6 address range, in CIDR notation.")
		ec2_provisionByoipCidrCmd.Flags().String("cidr-authorization-context", "", "A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP.")
		ec2_provisionByoipCidrCmd.Flags().String("description", "", "A description for the address range and the address pool.")
		ec2_provisionByoipCidrCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_provisionByoipCidrCmd.Flags().Bool("multi-region", false, "Reserved.")
		ec2_provisionByoipCidrCmd.Flags().String("network-border-group", "", "If you have [Local Zones](https://docs.aws.amazon.com/local-zones/latest/ug/how-local-zones-work.html) enabled, you can choose a network border group for Local Zones when you provision and advertise a BYOIPv4 CIDR.")
		ec2_provisionByoipCidrCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_provisionByoipCidrCmd.Flags().Bool("no-multi-region", false, "Reserved.")
		ec2_provisionByoipCidrCmd.Flags().Bool("no-publicly-advertisable", false, "(IPv6 only) Indicate whether the address range will be publicly advertised to the internet.")
		ec2_provisionByoipCidrCmd.Flags().String("pool-tag-specifications", "", "The tags to apply to the address pool.")
		ec2_provisionByoipCidrCmd.Flags().Bool("publicly-advertisable", false, "(IPv6 only) Indicate whether the address range will be publicly advertised to the internet.")
		ec2_provisionByoipCidrCmd.MarkFlagRequired("cidr")
		ec2_provisionByoipCidrCmd.Flag("no-dry-run").Hidden = true
		ec2_provisionByoipCidrCmd.Flag("no-multi-region").Hidden = true
		ec2_provisionByoipCidrCmd.Flag("no-publicly-advertisable").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_provisionByoipCidrCmd)
}
