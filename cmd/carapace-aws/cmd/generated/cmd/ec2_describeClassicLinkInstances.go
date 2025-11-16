package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeClassicLinkInstancesCmd = &cobra.Command{
	Use:   "describe-classic-link-instances",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeClassicLinkInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeClassicLinkInstancesCmd).Standalone()

		ec2_describeClassicLinkInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeClassicLinkInstancesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeClassicLinkInstancesCmd.Flags().String("instance-ids", "", "The instance IDs.")
		ec2_describeClassicLinkInstancesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeClassicLinkInstancesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeClassicLinkInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeClassicLinkInstancesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeClassicLinkInstancesCmd)
}
