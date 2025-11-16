package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTrafficMirrorSessionCmd = &cobra.Command{
	Use:   "delete-traffic-mirror-session",
	Short: "Deletes the specified Traffic Mirror session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTrafficMirrorSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTrafficMirrorSessionCmd).Standalone()

		ec2_deleteTrafficMirrorSessionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorSessionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorSessionCmd.Flags().String("traffic-mirror-session-id", "", "The ID of the Traffic Mirror session.")
		ec2_deleteTrafficMirrorSessionCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTrafficMirrorSessionCmd.MarkFlagRequired("traffic-mirror-session-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTrafficMirrorSessionCmd)
}
