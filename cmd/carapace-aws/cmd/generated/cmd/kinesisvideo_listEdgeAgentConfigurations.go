package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_listEdgeAgentConfigurationsCmd = &cobra.Command{
	Use:   "list-edge-agent-configurations",
	Short: "Returns an array of edge configurations associated with the specified Edge Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_listEdgeAgentConfigurationsCmd).Standalone()

	kinesisvideo_listEdgeAgentConfigurationsCmd.Flags().String("hub-device-arn", "", "The \"Internet of Things (IoT) Thing\" Arn of the edge agent.")
	kinesisvideo_listEdgeAgentConfigurationsCmd.Flags().String("max-results", "", "The maximum number of edge configurations to return in the response.")
	kinesisvideo_listEdgeAgentConfigurationsCmd.Flags().String("next-token", "", "If you specify this parameter, when the result of a `ListEdgeAgentConfigurations` operation is truncated, the call returns the `NextToken` in the response.")
	kinesisvideo_listEdgeAgentConfigurationsCmd.MarkFlagRequired("hub-device-arn")
	kinesisvideoCmd.AddCommand(kinesisvideo_listEdgeAgentConfigurationsCmd)
}
