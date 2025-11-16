package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFleetInstancesCmd = &cobra.Command{
	Use:   "describe-fleet-instances",
	Short: "Describes the running instances for the specified EC2 Fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFleetInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeFleetInstancesCmd).Standalone()

		ec2_describeFleetInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeFleetInstancesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeFleetInstancesCmd.Flags().String("fleet-id", "", "The ID of the EC2 Fleet.")
		ec2_describeFleetInstancesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeFleetInstancesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeFleetInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeFleetInstancesCmd.MarkFlagRequired("fleet-id")
		ec2_describeFleetInstancesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeFleetInstancesCmd)
}
