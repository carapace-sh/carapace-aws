package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateEndpointGroupCmd = &cobra.Command{
	Use:   "update-endpoint-group",
	Short: "Update an endpoint group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateEndpointGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_updateEndpointGroupCmd).Standalone()

		globalaccelerator_updateEndpointGroupCmd.Flags().String("endpoint-configurations", "", "The list of endpoint objects.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("health-check-interval-seconds", "", "The time—10 seconds or 30 seconds—between each health check for an endpoint.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("health-check-path", "", "If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("health-check-port", "", "The port that Global Accelerator uses to check the health of endpoints that are part of this endpoint group.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("health-check-protocol", "", "The protocol that Global Accelerator uses to check the health of endpoints that are part of this endpoint group.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("port-overrides", "", "Override specific listener ports used to route traffic to endpoints that are part of this endpoint group.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("threshold-count", "", "The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.")
		globalaccelerator_updateEndpointGroupCmd.Flags().String("traffic-dial-percentage", "", "The percentage of traffic to send to an Amazon Web Services Region.")
		globalaccelerator_updateEndpointGroupCmd.MarkFlagRequired("endpoint-group-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_updateEndpointGroupCmd)
}
