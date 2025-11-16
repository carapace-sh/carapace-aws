package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteSecurityGroupCmd = &cobra.Command{
	Use:   "delete-security-group",
	Short: "Deletes a security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteSecurityGroupCmd).Standalone()

	ec2_deleteSecurityGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteSecurityGroupCmd.Flags().String("group-id", "", "The ID of the security group.")
	ec2_deleteSecurityGroupCmd.Flags().String("group-name", "", "\\[Default VPC] The name of the security group.")
	ec2_deleteSecurityGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteSecurityGroupCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteSecurityGroupCmd)
}
