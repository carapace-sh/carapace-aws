package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSecurityGroupRulesCmd = &cobra.Command{
	Use:   "describe-security-group-rules",
	Short: "Describes one or more of your security group rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSecurityGroupRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSecurityGroupRulesCmd).Standalone()

		ec2_describeSecurityGroupRulesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSecurityGroupRulesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeSecurityGroupRulesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeSecurityGroupRulesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeSecurityGroupRulesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSecurityGroupRulesCmd.Flags().String("security-group-rule-ids", "", "The IDs of the security group rules.")
		ec2_describeSecurityGroupRulesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeSecurityGroupRulesCmd)
}
