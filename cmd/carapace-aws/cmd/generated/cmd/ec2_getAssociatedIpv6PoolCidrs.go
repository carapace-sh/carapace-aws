package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getAssociatedIpv6PoolCidrsCmd = &cobra.Command{
	Use:   "get-associated-ipv6-pool-cidrs",
	Short: "Gets information about the IPv6 CIDR block associations for a specified IPv6 address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getAssociatedIpv6PoolCidrsCmd).Standalone()

	ec2_getAssociatedIpv6PoolCidrsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getAssociatedIpv6PoolCidrsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_getAssociatedIpv6PoolCidrsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getAssociatedIpv6PoolCidrsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getAssociatedIpv6PoolCidrsCmd.Flags().String("pool-id", "", "The ID of the IPv6 address pool.")
	ec2_getAssociatedIpv6PoolCidrsCmd.Flag("no-dry-run").Hidden = true
	ec2_getAssociatedIpv6PoolCidrsCmd.MarkFlagRequired("pool-id")
	ec2Cmd.AddCommand(ec2_getAssociatedIpv6PoolCidrsCmd)
}
