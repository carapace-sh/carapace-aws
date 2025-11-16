package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTrafficMirrorSessionsCmd = &cobra.Command{
	Use:   "describe-traffic-mirror-sessions",
	Short: "Describes one or more Traffic Mirror sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTrafficMirrorSessionsCmd).Standalone()

	ec2_describeTrafficMirrorSessionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTrafficMirrorSessionsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeTrafficMirrorSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeTrafficMirrorSessionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeTrafficMirrorSessionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTrafficMirrorSessionsCmd.Flags().String("traffic-mirror-session-ids", "", "The ID of the Traffic Mirror session.")
	ec2_describeTrafficMirrorSessionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeTrafficMirrorSessionsCmd)
}
