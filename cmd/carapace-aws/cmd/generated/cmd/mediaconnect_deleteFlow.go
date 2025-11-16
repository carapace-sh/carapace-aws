package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_deleteFlowCmd = &cobra.Command{
	Use:   "delete-flow",
	Short: "Deletes a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_deleteFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_deleteFlowCmd).Standalone()

		mediaconnect_deleteFlowCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to delete.")
		mediaconnect_deleteFlowCmd.MarkFlagRequired("flow-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_deleteFlowCmd)
}
