package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifySecurityGroupRulesCmd = &cobra.Command{
	Use:   "modify-security-group-rules",
	Short: "Modifies the rules of a security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifySecurityGroupRulesCmd).Standalone()

	ec2_modifySecurityGroupRulesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifySecurityGroupRulesCmd.Flags().String("group-id", "", "The ID of the security group.")
	ec2_modifySecurityGroupRulesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifySecurityGroupRulesCmd.Flags().String("security-group-rules", "", "Information about the security group properties to update.")
	ec2_modifySecurityGroupRulesCmd.MarkFlagRequired("group-id")
	ec2_modifySecurityGroupRulesCmd.Flag("no-dry-run").Hidden = true
	ec2_modifySecurityGroupRulesCmd.MarkFlagRequired("security-group-rules")
	ec2Cmd.AddCommand(ec2_modifySecurityGroupRulesCmd)
}
