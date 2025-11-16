package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSecurityGroupsCmd = &cobra.Command{
	Use:   "describe-security-groups",
	Short: "Describes the specified security groups or all of your security groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSecurityGroupsCmd).Standalone()

	ec2_describeSecurityGroupsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSecurityGroupsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeSecurityGroupsCmd.Flags().String("group-ids", "", "The IDs of the security groups.")
	ec2_describeSecurityGroupsCmd.Flags().String("group-names", "", "\\[Default VPC] The names of the security groups.")
	ec2_describeSecurityGroupsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeSecurityGroupsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeSecurityGroupsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSecurityGroupsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeSecurityGroupsCmd)
}
