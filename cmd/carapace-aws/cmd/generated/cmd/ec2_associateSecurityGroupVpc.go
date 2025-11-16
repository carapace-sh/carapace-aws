package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateSecurityGroupVpcCmd = &cobra.Command{
	Use:   "associate-security-group-vpc",
	Short: "Associates a security group with another VPC in the same Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateSecurityGroupVpcCmd).Standalone()

	ec2_associateSecurityGroupVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateSecurityGroupVpcCmd.Flags().String("group-id", "", "A security group ID.")
	ec2_associateSecurityGroupVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateSecurityGroupVpcCmd.Flags().String("vpc-id", "", "A VPC ID.")
	ec2_associateSecurityGroupVpcCmd.MarkFlagRequired("group-id")
	ec2_associateSecurityGroupVpcCmd.Flag("no-dry-run").Hidden = true
	ec2_associateSecurityGroupVpcCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_associateSecurityGroupVpcCmd)
}
