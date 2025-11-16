package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTrafficMirrorSessionCmd = &cobra.Command{
	Use:   "create-traffic-mirror-session",
	Short: "Creates a Traffic Mirror session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTrafficMirrorSessionCmd).Standalone()

	ec2_createTrafficMirrorSessionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("description", "", "The description of the Traffic Mirror session.")
	ec2_createTrafficMirrorSessionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("network-interface-id", "", "The ID of the source network interface.")
	ec2_createTrafficMirrorSessionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("packet-length", "", "The number of bytes in each packet to mirror.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("session-number", "", "The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("tag-specifications", "", "The tags to assign to a Traffic Mirror session.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("traffic-mirror-filter-id", "", "The ID of the Traffic Mirror filter.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("traffic-mirror-target-id", "", "The ID of the Traffic Mirror target.")
	ec2_createTrafficMirrorSessionCmd.Flags().String("virtual-network-id", "", "The VXLAN ID for the Traffic Mirror session.")
	ec2_createTrafficMirrorSessionCmd.MarkFlagRequired("network-interface-id")
	ec2_createTrafficMirrorSessionCmd.Flag("no-dry-run").Hidden = true
	ec2_createTrafficMirrorSessionCmd.MarkFlagRequired("session-number")
	ec2_createTrafficMirrorSessionCmd.MarkFlagRequired("traffic-mirror-filter-id")
	ec2_createTrafficMirrorSessionCmd.MarkFlagRequired("traffic-mirror-target-id")
	ec2Cmd.AddCommand(ec2_createTrafficMirrorSessionCmd)
}
