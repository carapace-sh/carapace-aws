package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_revokeSecurityGroupEgressCmd = &cobra.Command{
	Use:   "revoke-security-group-egress",
	Short: "Removes the specified outbound (egress) rules from the specified security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_revokeSecurityGroupEgressCmd).Standalone()

	ec2_revokeSecurityGroupEgressCmd.Flags().String("cidr-ip", "", "Not supported.")
	ec2_revokeSecurityGroupEgressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("from-port", "", "Not supported.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("group-id", "", "The ID of the security group.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("ip-permissions", "", "The sets of IP permissions.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("ip-protocol", "", "Not supported.")
	ec2_revokeSecurityGroupEgressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("security-group-rule-ids", "", "The IDs of the security group rules.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("source-security-group-name", "", "Not supported.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("source-security-group-owner-id", "", "Not supported.")
	ec2_revokeSecurityGroupEgressCmd.Flags().String("to-port", "", "Not supported.")
	ec2_revokeSecurityGroupEgressCmd.MarkFlagRequired("group-id")
	ec2_revokeSecurityGroupEgressCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_revokeSecurityGroupEgressCmd)
}
