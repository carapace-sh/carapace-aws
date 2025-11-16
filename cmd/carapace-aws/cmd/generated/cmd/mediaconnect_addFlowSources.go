package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_addFlowSourcesCmd = &cobra.Command{
	Use:   "add-flow-sources",
	Short: "Adds sources to a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_addFlowSourcesCmd).Standalone()

	mediaconnect_addFlowSourcesCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to update.")
	mediaconnect_addFlowSourcesCmd.Flags().String("sources", "", "A list of sources that you want to add to the flow.")
	mediaconnect_addFlowSourcesCmd.MarkFlagRequired("flow-arn")
	mediaconnect_addFlowSourcesCmd.MarkFlagRequired("sources")
	mediaconnectCmd.AddCommand(mediaconnect_addFlowSourcesCmd)
}
