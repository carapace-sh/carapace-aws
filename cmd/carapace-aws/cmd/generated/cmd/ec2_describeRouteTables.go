package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeRouteTablesCmd = &cobra.Command{
	Use:   "describe-route-tables",
	Short: "Describes your route tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeRouteTablesCmd).Standalone()

	ec2_describeRouteTablesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeRouteTablesCmd.Flags().String("filters", "", "The filters.")
	ec2_describeRouteTablesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeRouteTablesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeRouteTablesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeRouteTablesCmd.Flags().String("route-table-ids", "", "The IDs of the route tables.")
	ec2_describeRouteTablesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeRouteTablesCmd)
}
