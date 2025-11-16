package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTrafficMirrorTargetsCmd = &cobra.Command{
	Use:   "describe-traffic-mirror-targets",
	Short: "Information about one or more Traffic Mirror targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTrafficMirrorTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTrafficMirrorTargetsCmd).Standalone()

		ec2_describeTrafficMirrorTargetsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTrafficMirrorTargetsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTrafficMirrorTargetsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTrafficMirrorTargetsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTrafficMirrorTargetsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTrafficMirrorTargetsCmd.Flags().String("traffic-mirror-target-ids", "", "The ID of the Traffic Mirror targets.")
		ec2_describeTrafficMirrorTargetsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTrafficMirrorTargetsCmd)
}
