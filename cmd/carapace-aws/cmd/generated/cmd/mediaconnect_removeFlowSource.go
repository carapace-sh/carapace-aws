package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_removeFlowSourceCmd = &cobra.Command{
	Use:   "remove-flow-source",
	Short: "Removes a source from an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_removeFlowSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_removeFlowSourceCmd).Standalone()

		mediaconnect_removeFlowSourceCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to remove a source from.")
		mediaconnect_removeFlowSourceCmd.Flags().String("source-arn", "", "The ARN of the source that you want to remove.")
		mediaconnect_removeFlowSourceCmd.MarkFlagRequired("flow-arn")
		mediaconnect_removeFlowSourceCmd.MarkFlagRequired("source-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_removeFlowSourceCmd)
}
