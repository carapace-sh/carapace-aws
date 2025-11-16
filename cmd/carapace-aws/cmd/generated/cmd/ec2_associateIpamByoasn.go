package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateIpamByoasnCmd = &cobra.Command{
	Use:   "associate-ipam-byoasn",
	Short: "Associates your Autonomous System Number (ASN) with a BYOIP CIDR that you own in the same Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateIpamByoasnCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_associateIpamByoasnCmd).Standalone()

		ec2_associateIpamByoasnCmd.Flags().String("asn", "", "A public 2-byte or 4-byte ASN.")
		ec2_associateIpamByoasnCmd.Flags().String("cidr", "", "The BYOIP CIDR you want to associate with an ASN.")
		ec2_associateIpamByoasnCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateIpamByoasnCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateIpamByoasnCmd.MarkFlagRequired("asn")
		ec2_associateIpamByoasnCmd.MarkFlagRequired("cidr")
		ec2_associateIpamByoasnCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_associateIpamByoasnCmd)
}
