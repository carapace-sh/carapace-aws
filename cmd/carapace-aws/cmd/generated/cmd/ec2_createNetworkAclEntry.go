package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createNetworkAclEntryCmd = &cobra.Command{
	Use:   "create-network-acl-entry",
	Short: "Creates an entry (a rule) in a network ACL with the specified rule number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createNetworkAclEntryCmd).Standalone()

	ec2_createNetworkAclEntryCmd.Flags().String("cidr-block", "", "The IPv4 network range to allow or deny, in CIDR notation (for example `172.16.0.0/24`). We modify the specified CIDR block to its canonical form; for example, if you specify `100.68.0.18/18`, we modify it to `100.68.0.0/18`.")
	ec2_createNetworkAclEntryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createNetworkAclEntryCmd.Flags().Bool("egress", false, "Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet).")
	ec2_createNetworkAclEntryCmd.Flags().String("icmp-type-code", "", "ICMP protocol: The ICMP or ICMPv6 type and code.")
	ec2_createNetworkAclEntryCmd.Flags().String("ipv6-cidr-block", "", "The IPv6 network range to allow or deny, in CIDR notation (for example `2001:db8:1234:1a00::/64`).")
	ec2_createNetworkAclEntryCmd.Flags().String("network-acl-id", "", "The ID of the network ACL.")
	ec2_createNetworkAclEntryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createNetworkAclEntryCmd.Flags().Bool("no-egress", false, "Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet).")
	ec2_createNetworkAclEntryCmd.Flags().String("port-range", "", "TCP or UDP protocols: The range of ports the rule applies to.")
	ec2_createNetworkAclEntryCmd.Flags().String("protocol", "", "The protocol number.")
	ec2_createNetworkAclEntryCmd.Flags().String("rule-action", "", "Indicates whether to allow or deny the traffic that matches the rule.")
	ec2_createNetworkAclEntryCmd.Flags().String("rule-number", "", "The rule number for the entry (for example, 100).")
	ec2_createNetworkAclEntryCmd.MarkFlagRequired("egress")
	ec2_createNetworkAclEntryCmd.MarkFlagRequired("network-acl-id")
	ec2_createNetworkAclEntryCmd.Flag("no-dry-run").Hidden = true
	ec2_createNetworkAclEntryCmd.Flag("no-egress").Hidden = true
	ec2_createNetworkAclEntryCmd.MarkFlagRequired("no-egress")
	ec2_createNetworkAclEntryCmd.MarkFlagRequired("protocol")
	ec2_createNetworkAclEntryCmd.MarkFlagRequired("rule-action")
	ec2_createNetworkAclEntryCmd.MarkFlagRequired("rule-number")
	ec2Cmd.AddCommand(ec2_createNetworkAclEntryCmd)
}
