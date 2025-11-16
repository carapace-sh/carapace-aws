package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_addFlowOutputsCmd = &cobra.Command{
	Use:   "add-flow-outputs",
	Short: "Adds outputs to an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_addFlowOutputsCmd).Standalone()

	mediaconnect_addFlowOutputsCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to add outputs to.")
	mediaconnect_addFlowOutputsCmd.Flags().String("outputs", "", "A list of outputs that you want to add to the flow.")
	mediaconnect_addFlowOutputsCmd.MarkFlagRequired("flow-arn")
	mediaconnect_addFlowOutputsCmd.MarkFlagRequired("outputs")
	mediaconnectCmd.AddCommand(mediaconnect_addFlowOutputsCmd)
}
