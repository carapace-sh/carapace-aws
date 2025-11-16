package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_startFlowCmd = &cobra.Command{
	Use:   "start-flow",
	Short: "Starts a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_startFlowCmd).Standalone()

	mediaconnect_startFlowCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to start.")
	mediaconnect_startFlowCmd.MarkFlagRequired("flow-arn")
	mediaconnectCmd.AddCommand(mediaconnect_startFlowCmd)
}
