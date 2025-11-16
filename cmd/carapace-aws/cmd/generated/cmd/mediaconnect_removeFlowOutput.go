package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_removeFlowOutputCmd = &cobra.Command{
	Use:   "remove-flow-output",
	Short: "Removes an output from an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_removeFlowOutputCmd).Standalone()

	mediaconnect_removeFlowOutputCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to remove an output from.")
	mediaconnect_removeFlowOutputCmd.Flags().String("output-arn", "", "The ARN of the output that you want to remove.")
	mediaconnect_removeFlowOutputCmd.MarkFlagRequired("flow-arn")
	mediaconnect_removeFlowOutputCmd.MarkFlagRequired("output-arn")
	mediaconnectCmd.AddCommand(mediaconnect_removeFlowOutputCmd)
}
