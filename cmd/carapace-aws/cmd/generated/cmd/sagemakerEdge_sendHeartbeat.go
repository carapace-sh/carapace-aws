package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerEdge_sendHeartbeatCmd = &cobra.Command{
	Use:   "send-heartbeat",
	Short: "Use to get the current status of devices registered on SageMaker Edge Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerEdge_sendHeartbeatCmd).Standalone()

	sagemakerEdge_sendHeartbeatCmd.Flags().String("agent-metrics", "", "For internal use.")
	sagemakerEdge_sendHeartbeatCmd.Flags().String("agent-version", "", "Returns the version of the agent.")
	sagemakerEdge_sendHeartbeatCmd.Flags().String("deployment-result", "", "Returns the result of a deployment on the device.")
	sagemakerEdge_sendHeartbeatCmd.Flags().String("device-fleet-name", "", "The name of the fleet that the device belongs to.")
	sagemakerEdge_sendHeartbeatCmd.Flags().String("device-name", "", "The unique name of the device.")
	sagemakerEdge_sendHeartbeatCmd.Flags().String("models", "", "Returns a list of models deployed on the the device.")
	sagemakerEdge_sendHeartbeatCmd.MarkFlagRequired("agent-version")
	sagemakerEdge_sendHeartbeatCmd.MarkFlagRequired("device-fleet-name")
	sagemakerEdge_sendHeartbeatCmd.MarkFlagRequired("device-name")
	sagemakerEdgeCmd.AddCommand(sagemakerEdge_sendHeartbeatCmd)
}
