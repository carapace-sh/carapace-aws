package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamByoasnCmd = &cobra.Command{
	Use:   "describe-ipam-byoasn",
	Short: "Describes your Autonomous System Numbers (ASNs), their provisioning statuses, and the BYOIP CIDRs with which they are associated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamByoasnCmd).Standalone()

	ec2_describeIpamByoasnCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeIpamByoasnCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeIpamByoasnCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeIpamByoasnCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeIpamByoasnCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeIpamByoasnCmd)
}
