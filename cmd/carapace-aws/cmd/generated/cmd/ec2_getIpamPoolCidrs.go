package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamPoolCidrsCmd = &cobra.Command{
	Use:   "get-ipam-pool-cidrs",
	Short: "Get the CIDRs provisioned to an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamPoolCidrsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getIpamPoolCidrsCmd).Standalone()

		ec2_getIpamPoolCidrsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamPoolCidrsCmd.Flags().String("filters", "", "One or more filters for the request.")
		ec2_getIpamPoolCidrsCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool you want the CIDR for.")
		ec2_getIpamPoolCidrsCmd.Flags().String("max-results", "", "The maximum number of results to return in the request.")
		ec2_getIpamPoolCidrsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getIpamPoolCidrsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamPoolCidrsCmd.MarkFlagRequired("ipam-pool-id")
		ec2_getIpamPoolCidrsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getIpamPoolCidrsCmd)
}
