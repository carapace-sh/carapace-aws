package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTrafficMirrorFilterCmd = &cobra.Command{
	Use:   "create-traffic-mirror-filter",
	Short: "Creates a Traffic Mirror filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTrafficMirrorFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTrafficMirrorFilterCmd).Standalone()

		ec2_createTrafficMirrorFilterCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createTrafficMirrorFilterCmd.Flags().String("description", "", "The description of the Traffic Mirror filter.")
		ec2_createTrafficMirrorFilterCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTrafficMirrorFilterCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTrafficMirrorFilterCmd.Flags().String("tag-specifications", "", "The tags to assign to a Traffic Mirror filter.")
		ec2_createTrafficMirrorFilterCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createTrafficMirrorFilterCmd)
}
