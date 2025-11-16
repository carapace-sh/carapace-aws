package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSpotInstanceRequestsCmd = &cobra.Command{
	Use:   "describe-spot-instance-requests",
	Short: "Describes the specified Spot Instance requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSpotInstanceRequestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSpotInstanceRequestsCmd).Standalone()

		ec2_describeSpotInstanceRequestsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotInstanceRequestsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeSpotInstanceRequestsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeSpotInstanceRequestsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeSpotInstanceRequestsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSpotInstanceRequestsCmd.Flags().String("spot-instance-request-ids", "", "The IDs of the Spot Instance requests.")
		ec2_describeSpotInstanceRequestsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeSpotInstanceRequestsCmd)
}
