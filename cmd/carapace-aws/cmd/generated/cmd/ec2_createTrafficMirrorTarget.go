package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTrafficMirrorTargetCmd = &cobra.Command{
	Use:   "create-traffic-mirror-target",
	Short: "Creates a target for your Traffic Mirror session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTrafficMirrorTargetCmd).Standalone()

	ec2_createTrafficMirrorTargetCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createTrafficMirrorTargetCmd.Flags().String("description", "", "The description of the Traffic Mirror target.")
	ec2_createTrafficMirrorTargetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTrafficMirrorTargetCmd.Flags().String("gateway-load-balancer-endpoint-id", "", "The ID of the Gateway Load Balancer endpoint.")
	ec2_createTrafficMirrorTargetCmd.Flags().String("network-interface-id", "", "The network interface ID that is associated with the target.")
	ec2_createTrafficMirrorTargetCmd.Flags().String("network-load-balancer-arn", "", "The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.")
	ec2_createTrafficMirrorTargetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTrafficMirrorTargetCmd.Flags().String("tag-specifications", "", "The tags to assign to the Traffic Mirror target.")
	ec2_createTrafficMirrorTargetCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createTrafficMirrorTargetCmd)
}
