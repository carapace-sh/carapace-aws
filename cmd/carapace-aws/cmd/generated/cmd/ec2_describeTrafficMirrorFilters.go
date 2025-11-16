package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTrafficMirrorFiltersCmd = &cobra.Command{
	Use:   "describe-traffic-mirror-filters",
	Short: "Describes one or more Traffic Mirror filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTrafficMirrorFiltersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTrafficMirrorFiltersCmd).Standalone()

		ec2_describeTrafficMirrorFiltersCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTrafficMirrorFiltersCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTrafficMirrorFiltersCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTrafficMirrorFiltersCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTrafficMirrorFiltersCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTrafficMirrorFiltersCmd.Flags().String("traffic-mirror-filter-ids", "", "The ID of the Traffic Mirror filter.")
		ec2_describeTrafficMirrorFiltersCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTrafficMirrorFiltersCmd)
}
