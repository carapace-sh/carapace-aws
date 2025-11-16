package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_createEndpointGroupCmd = &cobra.Command{
	Use:   "create-endpoint-group",
	Short: "Create an endpoint group for the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_createEndpointGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_createEndpointGroupCmd).Standalone()

		globalaccelerator_createEndpointGroupCmd.Flags().String("endpoint-configurations", "", "The list of endpoint objects.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("endpoint-group-region", "", "The Amazon Web Services Region where the endpoint group is located.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("health-check-interval-seconds", "", "The time—10 seconds or 30 seconds—between each health check for an endpoint.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("health-check-path", "", "If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("health-check-port", "", "The port that Global Accelerator uses to check the health of endpoints that are part of this endpoint group.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("health-check-protocol", "", "The protocol that Global Accelerator uses to check the health of endpoints that are part of this endpoint group.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("idempotency-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency—that is, the uniqueness—of the request.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("port-overrides", "", "Override specific listener ports used to route traffic to endpoints that are part of this endpoint group.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("threshold-count", "", "The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.")
		globalaccelerator_createEndpointGroupCmd.Flags().String("traffic-dial-percentage", "", "The percentage of traffic to send to an Amazon Web Services Region.")
		globalaccelerator_createEndpointGroupCmd.MarkFlagRequired("endpoint-group-region")
		globalaccelerator_createEndpointGroupCmd.MarkFlagRequired("idempotency-token")
		globalaccelerator_createEndpointGroupCmd.MarkFlagRequired("listener-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_createEndpointGroupCmd)
}
