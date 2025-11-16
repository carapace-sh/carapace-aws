package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkAclsCmd = &cobra.Command{
	Use:   "describe-network-acls",
	Short: "Describes your network ACLs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkAclsCmd).Standalone()

	ec2_describeNetworkAclsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkAclsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeNetworkAclsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeNetworkAclsCmd.Flags().String("network-acl-ids", "", "The IDs of the network ACLs.")
	ec2_describeNetworkAclsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeNetworkAclsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkAclsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeNetworkAclsCmd)
}
