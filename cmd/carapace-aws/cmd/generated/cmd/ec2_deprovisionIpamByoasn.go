package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deprovisionIpamByoasnCmd = &cobra.Command{
	Use:   "deprovision-ipam-byoasn",
	Short: "Deprovisions your Autonomous System Number (ASN) from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deprovisionIpamByoasnCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deprovisionIpamByoasnCmd).Standalone()

		ec2_deprovisionIpamByoasnCmd.Flags().String("asn", "", "An ASN.")
		ec2_deprovisionIpamByoasnCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deprovisionIpamByoasnCmd.Flags().String("ipam-id", "", "The IPAM ID.")
		ec2_deprovisionIpamByoasnCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deprovisionIpamByoasnCmd.MarkFlagRequired("asn")
		ec2_deprovisionIpamByoasnCmd.MarkFlagRequired("ipam-id")
		ec2_deprovisionIpamByoasnCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deprovisionIpamByoasnCmd)
}
