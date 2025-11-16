package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_createTargetGroupCmd = &cobra.Command{
	Use:   "create-target-group",
	Short: "Creates a target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_createTargetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_createTargetGroupCmd).Standalone()

		elbv2_createTargetGroupCmd.Flags().String("health-check-enabled", "", "Indicates whether health checks are enabled.")
		elbv2_createTargetGroupCmd.Flags().String("health-check-interval-seconds", "", "The approximate amount of time, in seconds, between health checks of an individual target.")
		elbv2_createTargetGroupCmd.Flags().String("health-check-path", "", "\\[HTTP/HTTPS health checks] The destination for health checks on the targets.")
		elbv2_createTargetGroupCmd.Flags().String("health-check-port", "", "The port the load balancer uses when performing health checks on targets.")
		elbv2_createTargetGroupCmd.Flags().String("health-check-protocol", "", "The protocol the load balancer uses when performing health checks on targets.")
		elbv2_createTargetGroupCmd.Flags().String("health-check-timeout-seconds", "", "The amount of time, in seconds, during which no response from a target means a failed health check.")
		elbv2_createTargetGroupCmd.Flags().String("healthy-threshold-count", "", "The number of consecutive health check successes required before considering a target healthy.")
		elbv2_createTargetGroupCmd.Flags().String("ip-address-type", "", "The IP address type.")
		elbv2_createTargetGroupCmd.Flags().String("matcher", "", "\\[HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.")
		elbv2_createTargetGroupCmd.Flags().String("name", "", "The name of the target group.")
		elbv2_createTargetGroupCmd.Flags().String("port", "", "The port on which the targets receive traffic.")
		elbv2_createTargetGroupCmd.Flags().String("protocol", "", "The protocol to use for routing traffic to the targets.")
		elbv2_createTargetGroupCmd.Flags().String("protocol-version", "", "\\[HTTP/HTTPS protocol] The protocol version.")
		elbv2_createTargetGroupCmd.Flags().String("tags", "", "The tags to assign to the target group.")
		elbv2_createTargetGroupCmd.Flags().String("target-type", "", "The type of target that you must specify when registering targets with this target group.")
		elbv2_createTargetGroupCmd.Flags().String("unhealthy-threshold-count", "", "The number of consecutive health check failures required before considering a target unhealthy.")
		elbv2_createTargetGroupCmd.Flags().String("vpc-id", "", "The identifier of the virtual private cloud (VPC).")
		elbv2_createTargetGroupCmd.MarkFlagRequired("name")
	})
	elbv2Cmd.AddCommand(elbv2_createTargetGroupCmd)
}
