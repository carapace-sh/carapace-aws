package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTrafficMirrorTargetCmd = &cobra.Command{
	Use:   "delete-traffic-mirror-target",
	Short: "Deletes the specified Traffic Mirror target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTrafficMirrorTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTrafficMirrorTargetCmd).Standalone()

		ec2_deleteTrafficMirrorTargetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorTargetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorTargetCmd.Flags().String("traffic-mirror-target-id", "", "The ID of the Traffic Mirror target.")
		ec2_deleteTrafficMirrorTargetCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTrafficMirrorTargetCmd.MarkFlagRequired("traffic-mirror-target-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTrafficMirrorTargetCmd)
}
