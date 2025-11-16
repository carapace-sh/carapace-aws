package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_updateSecurityGroupRuleDescriptionsEgressCmd = &cobra.Command{
	Use:   "update-security-group-rule-descriptions-egress",
	Short: "Updates the description of an egress (outbound) security group rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_updateSecurityGroupRuleDescriptionsEgressCmd).Standalone()

	ec2_updateSecurityGroupRuleDescriptionsEgressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_updateSecurityGroupRuleDescriptionsEgressCmd.Flags().String("group-id", "", "The ID of the security group.")
	ec2_updateSecurityGroupRuleDescriptionsEgressCmd.Flags().String("group-name", "", "\\[Default VPC] The name of the security group.")
	ec2_updateSecurityGroupRuleDescriptionsEgressCmd.Flags().String("ip-permissions", "", "The IP permissions for the security group rule.")
	ec2_updateSecurityGroupRuleDescriptionsEgressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_updateSecurityGroupRuleDescriptionsEgressCmd.Flags().String("security-group-rule-descriptions", "", "The description for the egress security group rules.")
	ec2_updateSecurityGroupRuleDescriptionsEgressCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_updateSecurityGroupRuleDescriptionsEgressCmd)
}
