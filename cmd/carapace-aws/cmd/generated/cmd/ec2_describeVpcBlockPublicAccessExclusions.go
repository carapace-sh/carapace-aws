package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcBlockPublicAccessExclusionsCmd = &cobra.Command{
	Use:   "describe-vpc-block-public-access-exclusions",
	Short: "Describe VPC Block Public Access (BPA) exclusions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcBlockPublicAccessExclusionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcBlockPublicAccessExclusionsCmd).Standalone()

		ec2_describeVpcBlockPublicAccessExclusionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcBlockPublicAccessExclusionsCmd.Flags().String("exclusion-ids", "", "IDs of exclusions.")
		ec2_describeVpcBlockPublicAccessExclusionsCmd.Flags().String("filters", "", "Filters for the request:")
		ec2_describeVpcBlockPublicAccessExclusionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeVpcBlockPublicAccessExclusionsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeVpcBlockPublicAccessExclusionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcBlockPublicAccessExclusionsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeVpcBlockPublicAccessExclusionsCmd)
}
