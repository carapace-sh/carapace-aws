package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_stopFlowCmd = &cobra.Command{
	Use:   "stop-flow",
	Short: "Stops a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_stopFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_stopFlowCmd).Standalone()

		mediaconnect_stopFlowCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to stop.")
		mediaconnect_stopFlowCmd.MarkFlagRequired("flow-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_stopFlowCmd)
}
