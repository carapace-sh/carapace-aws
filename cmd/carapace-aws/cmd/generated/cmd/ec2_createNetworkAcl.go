package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createNetworkAclCmd = &cobra.Command{
	Use:   "create-network-acl",
	Short: "Creates a network ACL in a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createNetworkAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createNetworkAclCmd).Standalone()

		ec2_createNetworkAclCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createNetworkAclCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkAclCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkAclCmd.Flags().String("tag-specifications", "", "The tags to assign to the network ACL.")
		ec2_createNetworkAclCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_createNetworkAclCmd.Flag("no-dry-run").Hidden = true
		ec2_createNetworkAclCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_createNetworkAclCmd)
}
