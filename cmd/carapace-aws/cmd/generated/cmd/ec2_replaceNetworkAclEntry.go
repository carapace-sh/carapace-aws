package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceNetworkAclEntryCmd = &cobra.Command{
	Use:   "replace-network-acl-entry",
	Short: "Replaces an entry (rule) in a network ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceNetworkAclEntryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_replaceNetworkAclEntryCmd).Standalone()

		ec2_replaceNetworkAclEntryCmd.Flags().String("cidr-block", "", "The IPv4 network range to allow or deny, in CIDR notation (for example `172.16.0.0/24`).")
		ec2_replaceNetworkAclEntryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_replaceNetworkAclEntryCmd.Flags().Bool("egress", false, "Indicates whether to replace the egress rule.")
		ec2_replaceNetworkAclEntryCmd.Flags().String("icmp-type-code", "", "ICMP protocol: The ICMP or ICMPv6 type and code.")
		ec2_replaceNetworkAclEntryCmd.Flags().String("ipv6-cidr-block", "", "The IPv6 network range to allow or deny, in CIDR notation (for example `2001:bd8:1234:1a00::/64`).")
		ec2_replaceNetworkAclEntryCmd.Flags().String("network-acl-id", "", "The ID of the ACL.")
		ec2_replaceNetworkAclEntryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_replaceNetworkAclEntryCmd.Flags().Bool("no-egress", false, "Indicates whether to replace the egress rule.")
		ec2_replaceNetworkAclEntryCmd.Flags().String("port-range", "", "TCP or UDP protocols: The range of ports the rule applies to.")
		ec2_replaceNetworkAclEntryCmd.Flags().String("protocol", "", "The protocol number.")
		ec2_replaceNetworkAclEntryCmd.Flags().String("rule-action", "", "Indicates whether to allow or deny the traffic that matches the rule.")
		ec2_replaceNetworkAclEntryCmd.Flags().String("rule-number", "", "The rule number of the entry to replace.")
		ec2_replaceNetworkAclEntryCmd.MarkFlagRequired("egress")
		ec2_replaceNetworkAclEntryCmd.MarkFlagRequired("network-acl-id")
		ec2_replaceNetworkAclEntryCmd.Flag("no-dry-run").Hidden = true
		ec2_replaceNetworkAclEntryCmd.Flag("no-egress").Hidden = true
		ec2_replaceNetworkAclEntryCmd.MarkFlagRequired("no-egress")
		ec2_replaceNetworkAclEntryCmd.MarkFlagRequired("protocol")
		ec2_replaceNetworkAclEntryCmd.MarkFlagRequired("rule-action")
		ec2_replaceNetworkAclEntryCmd.MarkFlagRequired("rule-number")
	})
	ec2Cmd.AddCommand(ec2_replaceNetworkAclEntryCmd)
}
