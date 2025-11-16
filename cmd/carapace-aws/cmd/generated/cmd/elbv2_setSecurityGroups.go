package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_setSecurityGroupsCmd = &cobra.Command{
	Use:   "set-security-groups",
	Short: "Associates the specified security groups with the specified Application Load Balancer or Network Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_setSecurityGroupsCmd).Standalone()

	elbv2_setSecurityGroupsCmd.Flags().String("enforce-security-group-inbound-rules-on-private-link-traffic", "", "Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through Amazon Web Services PrivateLink.")
	elbv2_setSecurityGroupsCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
	elbv2_setSecurityGroupsCmd.Flags().String("security-groups", "", "The IDs of the security groups.")
	elbv2_setSecurityGroupsCmd.MarkFlagRequired("load-balancer-arn")
	elbv2_setSecurityGroupsCmd.MarkFlagRequired("security-groups")
	elbv2Cmd.AddCommand(elbv2_setSecurityGroupsCmd)
}
