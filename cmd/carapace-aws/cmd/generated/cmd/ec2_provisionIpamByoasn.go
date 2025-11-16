package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_provisionIpamByoasnCmd = &cobra.Command{
	Use:   "provision-ipam-byoasn",
	Short: "Provisions your Autonomous System Number (ASN) for use in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_provisionIpamByoasnCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_provisionIpamByoasnCmd).Standalone()

		ec2_provisionIpamByoasnCmd.Flags().String("asn", "", "A public 2-byte or 4-byte ASN.")
		ec2_provisionIpamByoasnCmd.Flags().String("asn-authorization-context", "", "An ASN authorization context.")
		ec2_provisionIpamByoasnCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_provisionIpamByoasnCmd.Flags().String("ipam-id", "", "An IPAM ID.")
		ec2_provisionIpamByoasnCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_provisionIpamByoasnCmd.MarkFlagRequired("asn")
		ec2_provisionIpamByoasnCmd.MarkFlagRequired("asn-authorization-context")
		ec2_provisionIpamByoasnCmd.MarkFlagRequired("ipam-id")
		ec2_provisionIpamByoasnCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_provisionIpamByoasnCmd)
}
