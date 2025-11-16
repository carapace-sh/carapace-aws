package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNetworkAclCmd = &cobra.Command{
	Use:   "delete-network-acl",
	Short: "Deletes the specified network ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNetworkAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteNetworkAclCmd).Standalone()

		ec2_deleteNetworkAclCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteNetworkAclCmd.Flags().String("network-acl-id", "", "The ID of the network ACL.")
		ec2_deleteNetworkAclCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteNetworkAclCmd.MarkFlagRequired("network-acl-id")
		ec2_deleteNetworkAclCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteNetworkAclCmd)
}
