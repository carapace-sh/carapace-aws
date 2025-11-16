package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateFlowCmd = &cobra.Command{
	Use:   "update-flow",
	Short: "Updates an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateFlowCmd).Standalone()

	mediaconnect_updateFlowCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to update.")
	mediaconnect_updateFlowCmd.Flags().String("flow-size", "", "Determines the processing capacity and feature set of the flow.")
	mediaconnect_updateFlowCmd.Flags().String("maintenance", "", "The maintenance setting of the flow.")
	mediaconnect_updateFlowCmd.Flags().String("ndi-config", "", "Specifies the configuration settings for NDI outputs.")
	mediaconnect_updateFlowCmd.Flags().String("source-failover-config", "", "The settings for source failover.")
	mediaconnect_updateFlowCmd.Flags().String("source-monitoring-config", "", "The settings for source monitoring.")
	mediaconnect_updateFlowCmd.MarkFlagRequired("flow-arn")
	mediaconnectCmd.AddCommand(mediaconnect_updateFlowCmd)
}
