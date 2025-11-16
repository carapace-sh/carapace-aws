package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_createFlowCmd = &cobra.Command{
	Use:   "create-flow",
	Short: "Creates a new flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_createFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_createFlowCmd).Standalone()

		mediaconnect_createFlowCmd.Flags().String("availability-zone", "", "The Availability Zone that you want to create the flow in.")
		mediaconnect_createFlowCmd.Flags().String("entitlements", "", "The entitlements that you want to grant on a flow.")
		mediaconnect_createFlowCmd.Flags().String("flow-size", "", "Determines the processing capacity and feature set of the flow.")
		mediaconnect_createFlowCmd.Flags().String("flow-tags", "", "The key-value pairs that can be used to tag and organize the flow.")
		mediaconnect_createFlowCmd.Flags().String("maintenance", "", "The maintenance settings you want to use for the flow.")
		mediaconnect_createFlowCmd.Flags().String("media-streams", "", "The media streams that you want to add to the flow.")
		mediaconnect_createFlowCmd.Flags().String("name", "", "The name of the flow.")
		mediaconnect_createFlowCmd.Flags().String("ndi-config", "", "Specifies the configuration settings for NDI outputs.")
		mediaconnect_createFlowCmd.Flags().String("outputs", "", "The outputs that you want to add to this flow.")
		mediaconnect_createFlowCmd.Flags().String("source", "", "The settings for the source that you want to use for the new flow.")
		mediaconnect_createFlowCmd.Flags().String("source-failover-config", "", "The settings for source failover.")
		mediaconnect_createFlowCmd.Flags().String("source-monitoring-config", "", "The settings for source monitoring.")
		mediaconnect_createFlowCmd.Flags().String("sources", "", "The sources that are assigned to the flow.")
		mediaconnect_createFlowCmd.Flags().String("vpc-interfaces", "", "The VPC interfaces you want on the flow.")
		mediaconnect_createFlowCmd.MarkFlagRequired("name")
	})
	mediaconnectCmd.AddCommand(mediaconnect_createFlowCmd)
}
