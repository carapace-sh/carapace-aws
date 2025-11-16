package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateSecurityGroupVpcCmd = &cobra.Command{
	Use:   "disassociate-security-group-vpc",
	Short: "Disassociates a security group from a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateSecurityGroupVpcCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateSecurityGroupVpcCmd).Standalone()

		ec2_disassociateSecurityGroupVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateSecurityGroupVpcCmd.Flags().String("group-id", "", "A security group ID.")
		ec2_disassociateSecurityGroupVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateSecurityGroupVpcCmd.Flags().String("vpc-id", "", "A VPC ID.")
		ec2_disassociateSecurityGroupVpcCmd.MarkFlagRequired("group-id")
		ec2_disassociateSecurityGroupVpcCmd.Flag("no-dry-run").Hidden = true
		ec2_disassociateSecurityGroupVpcCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_disassociateSecurityGroupVpcCmd)
}
