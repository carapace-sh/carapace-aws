package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamResourceCidrsCmd = &cobra.Command{
	Use:   "get-ipam-resource-cidrs",
	Short: "Returns resource CIDRs managed by IPAM in a given scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamResourceCidrsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getIpamResourceCidrsCmd).Standalone()

		ec2_getIpamResourceCidrsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamResourceCidrsCmd.Flags().String("filters", "", "One or more filters for the request.")
		ec2_getIpamResourceCidrsCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool that the resource is in.")
		ec2_getIpamResourceCidrsCmd.Flags().String("ipam-scope-id", "", "The ID of the scope that the resource is in.")
		ec2_getIpamResourceCidrsCmd.Flags().String("max-results", "", "The maximum number of results to return in the request.")
		ec2_getIpamResourceCidrsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getIpamResourceCidrsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamResourceCidrsCmd.Flags().String("resource-id", "", "The ID of the resource.")
		ec2_getIpamResourceCidrsCmd.Flags().String("resource-owner", "", "The ID of the Amazon Web Services account that owns the resource.")
		ec2_getIpamResourceCidrsCmd.Flags().String("resource-tag", "", "The resource tag.")
		ec2_getIpamResourceCidrsCmd.Flags().String("resource-type", "", "The resource type.")
		ec2_getIpamResourceCidrsCmd.MarkFlagRequired("ipam-scope-id")
		ec2_getIpamResourceCidrsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getIpamResourceCidrsCmd)
}
