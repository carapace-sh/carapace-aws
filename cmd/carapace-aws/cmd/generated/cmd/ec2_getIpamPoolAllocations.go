package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamPoolAllocationsCmd = &cobra.Command{
	Use:   "get-ipam-pool-allocations",
	Short: "Get a list of all the CIDR allocations in an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamPoolAllocationsCmd).Standalone()

	ec2_getIpamPoolAllocationsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamPoolAllocationsCmd.Flags().String("filters", "", "One or more filters for the request.")
	ec2_getIpamPoolAllocationsCmd.Flags().String("ipam-pool-allocation-id", "", "The ID of the allocation.")
	ec2_getIpamPoolAllocationsCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool you want to see the allocations for.")
	ec2_getIpamPoolAllocationsCmd.Flags().String("max-results", "", "The maximum number of results you would like returned per page.")
	ec2_getIpamPoolAllocationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getIpamPoolAllocationsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamPoolAllocationsCmd.MarkFlagRequired("ipam-pool-id")
	ec2_getIpamPoolAllocationsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getIpamPoolAllocationsCmd)
}
