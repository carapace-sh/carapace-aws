package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createSecurityGroupCmd = &cobra.Command{
	Use:   "create-security-group",
	Short: "Creates a security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createSecurityGroupCmd).Standalone()

	ec2_createSecurityGroupCmd.Flags().String("description", "", "A description for the security group.")
	ec2_createSecurityGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createSecurityGroupCmd.Flags().String("group-name", "", "The name of the security group.")
	ec2_createSecurityGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createSecurityGroupCmd.Flags().String("tag-specifications", "", "The tags to assign to the security group.")
	ec2_createSecurityGroupCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_createSecurityGroupCmd.MarkFlagRequired("description")
	ec2_createSecurityGroupCmd.MarkFlagRequired("group-name")
	ec2_createSecurityGroupCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createSecurityGroupCmd)
}
