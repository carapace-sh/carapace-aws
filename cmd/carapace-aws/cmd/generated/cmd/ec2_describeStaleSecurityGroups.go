package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeStaleSecurityGroupsCmd = &cobra.Command{
	Use:   "describe-stale-security-groups",
	Short: "Describes the stale security group rules for security groups referenced across a VPC peering connection, transit gateway connection, or with a security group VPC association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeStaleSecurityGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeStaleSecurityGroupsCmd).Standalone()

		ec2_describeStaleSecurityGroupsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeStaleSecurityGroupsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeStaleSecurityGroupsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeStaleSecurityGroupsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeStaleSecurityGroupsCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_describeStaleSecurityGroupsCmd.Flag("no-dry-run").Hidden = true
		ec2_describeStaleSecurityGroupsCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_describeStaleSecurityGroupsCmd)
}
