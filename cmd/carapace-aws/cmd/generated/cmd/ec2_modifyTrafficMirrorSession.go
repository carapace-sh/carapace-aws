package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyTrafficMirrorSessionCmd = &cobra.Command{
	Use:   "modify-traffic-mirror-session",
	Short: "Modifies a Traffic Mirror session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyTrafficMirrorSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyTrafficMirrorSessionCmd).Standalone()

		ec2_modifyTrafficMirrorSessionCmd.Flags().String("description", "", "The description to assign to the Traffic Mirror session.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().String("packet-length", "", "The number of bytes in each packet to mirror.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().String("remove-fields", "", "The properties that you want to remove from the Traffic Mirror session.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().String("session-number", "", "The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().String("traffic-mirror-filter-id", "", "The ID of the Traffic Mirror filter.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().String("traffic-mirror-session-id", "", "The ID of the Traffic Mirror session.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().String("traffic-mirror-target-id", "", "The Traffic Mirror target.")
		ec2_modifyTrafficMirrorSessionCmd.Flags().String("virtual-network-id", "", "The virtual network ID of the Traffic Mirror session.")
		ec2_modifyTrafficMirrorSessionCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyTrafficMirrorSessionCmd.MarkFlagRequired("traffic-mirror-session-id")
	})
	ec2Cmd.AddCommand(ec2_modifyTrafficMirrorSessionCmd)
}
