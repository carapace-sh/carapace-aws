package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteIpamPoolCmd = &cobra.Command{
	Use:   "delete-ipam-pool",
	Short: "Delete an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteIpamPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteIpamPoolCmd).Standalone()

		ec2_deleteIpamPoolCmd.Flags().Bool("cascade", false, "Enables you to quickly delete an IPAM pool and all resources within that pool, including provisioned CIDRs, allocations, and other pools.")
		ec2_deleteIpamPoolCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamPoolCmd.Flags().String("ipam-pool-id", "", "The ID of the pool to delete.")
		ec2_deleteIpamPoolCmd.Flags().Bool("no-cascade", false, "Enables you to quickly delete an IPAM pool and all resources within that pool, including provisioned CIDRs, allocations, and other pools.")
		ec2_deleteIpamPoolCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamPoolCmd.MarkFlagRequired("ipam-pool-id")
		ec2_deleteIpamPoolCmd.Flag("no-cascade").Hidden = true
		ec2_deleteIpamPoolCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteIpamPoolCmd)
}
