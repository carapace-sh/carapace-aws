package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceNetworkAclAssociationCmd = &cobra.Command{
	Use:   "replace-network-acl-association",
	Short: "Changes which network ACL a subnet is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceNetworkAclAssociationCmd).Standalone()

	ec2_replaceNetworkAclAssociationCmd.Flags().String("association-id", "", "The ID of the current association between the original network ACL and the subnet.")
	ec2_replaceNetworkAclAssociationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceNetworkAclAssociationCmd.Flags().String("network-acl-id", "", "The ID of the new network ACL to associate with the subnet.")
	ec2_replaceNetworkAclAssociationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceNetworkAclAssociationCmd.MarkFlagRequired("association-id")
	ec2_replaceNetworkAclAssociationCmd.MarkFlagRequired("network-acl-id")
	ec2_replaceNetworkAclAssociationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_replaceNetworkAclAssociationCmd)
}
