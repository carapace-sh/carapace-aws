package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateIpamByoasnCmd = &cobra.Command{
	Use:   "disassociate-ipam-byoasn",
	Short: "Remove the association between your Autonomous System Number (ASN) and your BYOIP CIDR.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateIpamByoasnCmd).Standalone()

	ec2_disassociateIpamByoasnCmd.Flags().String("asn", "", "A public 2-byte or 4-byte ASN.")
	ec2_disassociateIpamByoasnCmd.Flags().String("cidr", "", "A BYOIP CIDR.")
	ec2_disassociateIpamByoasnCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateIpamByoasnCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateIpamByoasnCmd.MarkFlagRequired("asn")
	ec2_disassociateIpamByoasnCmd.MarkFlagRequired("cidr")
	ec2_disassociateIpamByoasnCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disassociateIpamByoasnCmd)
}
