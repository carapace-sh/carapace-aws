package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteIpamCmd = &cobra.Command{
	Use:   "delete-ipam",
	Short: "Delete an IPAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteIpamCmd).Standalone()

	ec2_deleteIpamCmd.Flags().Bool("cascade", false, "Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.")
	ec2_deleteIpamCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteIpamCmd.Flags().String("ipam-id", "", "The ID of the IPAM to delete.")
	ec2_deleteIpamCmd.Flags().Bool("no-cascade", false, "Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.")
	ec2_deleteIpamCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deleteIpamCmd.MarkFlagRequired("ipam-id")
	ec2_deleteIpamCmd.Flag("no-cascade").Hidden = true
	ec2_deleteIpamCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteIpamCmd)
}
