package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_authorizeSecurityGroupEgressCmd = &cobra.Command{
	Use:   "authorize-security-group-egress",
	Short: "Adds the specified outbound (egress) rules to a security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_authorizeSecurityGroupEgressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_authorizeSecurityGroupEgressCmd).Standalone()

		ec2_authorizeSecurityGroupEgressCmd.Flags().String("cidr-ip", "", "Not supported.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("from-port", "", "Not supported.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("group-id", "", "The ID of the security group.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("ip-permissions", "", "The permissions for the security group rules.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("ip-protocol", "", "Not supported.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("source-security-group-name", "", "Not supported.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("source-security-group-owner-id", "", "Not supported.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("tag-specifications", "", "The tags applied to the security group rule.")
		ec2_authorizeSecurityGroupEgressCmd.Flags().String("to-port", "", "Not supported.")
		ec2_authorizeSecurityGroupEgressCmd.MarkFlagRequired("group-id")
		ec2_authorizeSecurityGroupEgressCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_authorizeSecurityGroupEgressCmd)
}
