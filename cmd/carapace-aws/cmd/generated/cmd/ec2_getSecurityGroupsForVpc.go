package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getSecurityGroupsForVpcCmd = &cobra.Command{
	Use:   "get-security-groups-for-vpc",
	Short: "Gets security groups that can be associated by the Amazon Web Services account making the request with network interfaces in the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getSecurityGroupsForVpcCmd).Standalone()

	ec2_getSecurityGroupsForVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSecurityGroupsForVpcCmd.Flags().String("filters", "", "The filters.")
	ec2_getSecurityGroupsForVpcCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_getSecurityGroupsForVpcCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_getSecurityGroupsForVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSecurityGroupsForVpcCmd.Flags().String("vpc-id", "", "The VPC ID where the security group can be used.")
	ec2_getSecurityGroupsForVpcCmd.Flag("no-dry-run").Hidden = true
	ec2_getSecurityGroupsForVpcCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_getSecurityGroupsForVpcCmd)
}
