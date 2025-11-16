package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyTargetGroupCmd = &cobra.Command{
	Use:   "modify-target-group",
	Short: "Modifies the health checks used when evaluating the health state of the targets in the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyTargetGroupCmd).Standalone()

	elbv2_modifyTargetGroupCmd.Flags().String("health-check-enabled", "", "Indicates whether health checks are enabled.")
	elbv2_modifyTargetGroupCmd.Flags().String("health-check-interval-seconds", "", "The approximate amount of time, in seconds, between health checks of an individual target.")
	elbv2_modifyTargetGroupCmd.Flags().String("health-check-path", "", "\\[HTTP/HTTPS health checks] The destination for health checks on the targets.")
	elbv2_modifyTargetGroupCmd.Flags().String("health-check-port", "", "The port the load balancer uses when performing health checks on targets.")
	elbv2_modifyTargetGroupCmd.Flags().String("health-check-protocol", "", "The protocol the load balancer uses when performing health checks on targets.")
	elbv2_modifyTargetGroupCmd.Flags().String("health-check-timeout-seconds", "", "\\[HTTP/HTTPS health checks] The amount of time, in seconds, during which no response means a failed health check.")
	elbv2_modifyTargetGroupCmd.Flags().String("healthy-threshold-count", "", "The number of consecutive health checks successes required before considering an unhealthy target healthy.")
	elbv2_modifyTargetGroupCmd.Flags().String("matcher", "", "\\[HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.")
	elbv2_modifyTargetGroupCmd.Flags().String("target-group-arn", "", "The Amazon Resource Name (ARN) of the target group.")
	elbv2_modifyTargetGroupCmd.Flags().String("unhealthy-threshold-count", "", "The number of consecutive health check failures required before considering the target unhealthy.")
	elbv2_modifyTargetGroupCmd.MarkFlagRequired("target-group-arn")
	elbv2Cmd.AddCommand(elbv2_modifyTargetGroupCmd)
}
