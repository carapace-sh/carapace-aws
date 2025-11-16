package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_revokeSecurityGroupIngressCmd = &cobra.Command{
	Use:   "revoke-security-group-ingress",
	Short: "Removes the specified inbound (ingress) rules from a security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_revokeSecurityGroupIngressCmd).Standalone()

	ec2_revokeSecurityGroupIngressCmd.Flags().String("cidr-ip", "", "The CIDR IP address range.")
	ec2_revokeSecurityGroupIngressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("from-port", "", "If the protocol is TCP or UDP, this is the start of the port range.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("group-id", "", "The ID of the security group.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("group-name", "", "\\[Default VPC] The name of the security group.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("ip-permissions", "", "The sets of IP permissions.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("ip-protocol", "", "The IP protocol name (`tcp`, `udp`, `icmp`) or number (see [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)). Use `-1` to specify all.")
	ec2_revokeSecurityGroupIngressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("security-group-rule-ids", "", "The IDs of the security group rules.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("source-security-group-name", "", "\\[Default VPC] The name of the source security group.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("source-security-group-owner-id", "", "Not supported.")
	ec2_revokeSecurityGroupIngressCmd.Flags().String("to-port", "", "If the protocol is TCP or UDP, this is the end of the port range.")
	ec2_revokeSecurityGroupIngressCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_revokeSecurityGroupIngressCmd)
}
