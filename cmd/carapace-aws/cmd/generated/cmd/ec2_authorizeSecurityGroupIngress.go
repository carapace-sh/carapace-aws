package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_authorizeSecurityGroupIngressCmd = &cobra.Command{
	Use:   "authorize-security-group-ingress",
	Short: "Adds the specified inbound (ingress) rules to a security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_authorizeSecurityGroupIngressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_authorizeSecurityGroupIngressCmd).Standalone()

		ec2_authorizeSecurityGroupIngressCmd.Flags().String("cidr-ip", "", "The IPv4 address range, in CIDR format.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("from-port", "", "If the protocol is TCP or UDP, this is the start of the port range.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("group-id", "", "The ID of the security group.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("group-name", "", "\\[Default VPC] The name of the security group.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("ip-permissions", "", "The permissions for the security group rules.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("ip-protocol", "", "The IP protocol name (`tcp`, `udp`, `icmp`) or number (see [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)). To specify all protocols, use `-1`.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("source-security-group-name", "", "\\[Default VPC] The name of the source security group.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("source-security-group-owner-id", "", "The Amazon Web Services account ID for the source security group, if the source security group is in a different account.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("tag-specifications", "", "The tags applied to the security group rule.")
		ec2_authorizeSecurityGroupIngressCmd.Flags().String("to-port", "", "If the protocol is TCP or UDP, this is the end of the port range.")
		ec2_authorizeSecurityGroupIngressCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_authorizeSecurityGroupIngressCmd)
}
