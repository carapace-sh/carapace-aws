package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNetworkAclEntryCmd = &cobra.Command{
	Use:   "delete-network-acl-entry",
	Short: "Deletes the specified ingress or egress entry (rule) from the specified network ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNetworkAclEntryCmd).Standalone()

	ec2_deleteNetworkAclEntryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkAclEntryCmd.Flags().Bool("egress", false, "Indicates whether the rule is an egress rule.")
	ec2_deleteNetworkAclEntryCmd.Flags().String("network-acl-id", "", "The ID of the network ACL.")
	ec2_deleteNetworkAclEntryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkAclEntryCmd.Flags().Bool("no-egress", false, "Indicates whether the rule is an egress rule.")
	ec2_deleteNetworkAclEntryCmd.Flags().String("rule-number", "", "The rule number of the entry to delete.")
	ec2_deleteNetworkAclEntryCmd.MarkFlagRequired("egress")
	ec2_deleteNetworkAclEntryCmd.MarkFlagRequired("network-acl-id")
	ec2_deleteNetworkAclEntryCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteNetworkAclEntryCmd.Flag("no-egress").Hidden = true
	ec2_deleteNetworkAclEntryCmd.MarkFlagRequired("no-egress")
	ec2_deleteNetworkAclEntryCmd.MarkFlagRequired("rule-number")
	ec2Cmd.AddCommand(ec2_deleteNetworkAclEntryCmd)
}
