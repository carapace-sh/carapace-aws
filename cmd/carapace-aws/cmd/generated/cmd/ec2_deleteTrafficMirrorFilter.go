package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTrafficMirrorFilterCmd = &cobra.Command{
	Use:   "delete-traffic-mirror-filter",
	Short: "Deletes the specified Traffic Mirror filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTrafficMirrorFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTrafficMirrorFilterCmd).Standalone()

		ec2_deleteTrafficMirrorFilterCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorFilterCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTrafficMirrorFilterCmd.Flags().String("traffic-mirror-filter-id", "", "The ID of the Traffic Mirror filter.")
		ec2_deleteTrafficMirrorFilterCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTrafficMirrorFilterCmd.MarkFlagRequired("traffic-mirror-filter-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTrafficMirrorFilterCmd)
}
