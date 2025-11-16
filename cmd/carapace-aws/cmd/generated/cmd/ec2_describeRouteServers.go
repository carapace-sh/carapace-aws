package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeRouteServersCmd = &cobra.Command{
	Use:   "describe-route-servers",
	Short: "Describes one or more route servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeRouteServersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeRouteServersCmd).Standalone()

		ec2_describeRouteServersCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeRouteServersCmd.Flags().String("filters", "", "One or more filters to apply to the describe request.")
		ec2_describeRouteServersCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeRouteServersCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeRouteServersCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeRouteServersCmd.Flags().String("route-server-ids", "", "The IDs of the route servers to describe.")
		ec2_describeRouteServersCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeRouteServersCmd)
}
