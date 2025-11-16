package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_updateSecurityGroupRuleDescriptionsIngressCmd = &cobra.Command{
	Use:   "update-security-group-rule-descriptions-ingress",
	Short: "Updates the description of an ingress (inbound) security group rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_updateSecurityGroupRuleDescriptionsIngressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_updateSecurityGroupRuleDescriptionsIngressCmd).Standalone()

		ec2_updateSecurityGroupRuleDescriptionsIngressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_updateSecurityGroupRuleDescriptionsIngressCmd.Flags().String("group-id", "", "The ID of the security group.")
		ec2_updateSecurityGroupRuleDescriptionsIngressCmd.Flags().String("group-name", "", "\\[Default VPC] The name of the security group.")
		ec2_updateSecurityGroupRuleDescriptionsIngressCmd.Flags().String("ip-permissions", "", "The IP permissions for the security group rule.")
		ec2_updateSecurityGroupRuleDescriptionsIngressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_updateSecurityGroupRuleDescriptionsIngressCmd.Flags().String("security-group-rule-descriptions", "", "The description for the ingress security group rules.")
		ec2_updateSecurityGroupRuleDescriptionsIngressCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_updateSecurityGroupRuleDescriptionsIngressCmd)
}
