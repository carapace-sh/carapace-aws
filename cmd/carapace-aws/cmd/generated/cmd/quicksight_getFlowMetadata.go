package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_getFlowMetadataCmd = &cobra.Command{
	Use:   "get-flow-metadata",
	Short: "Retrieves the metadata of a flow, not including its definition specifying the steps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_getFlowMetadataCmd).Standalone()

	quicksight_getFlowMetadataCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the flow that you are getting metadata for.")
	quicksight_getFlowMetadataCmd.Flags().String("flow-id", "", "The unique identifier of the flow.")
	quicksight_getFlowMetadataCmd.MarkFlagRequired("aws-account-id")
	quicksight_getFlowMetadataCmd.MarkFlagRequired("flow-id")
	quicksightCmd.AddCommand(quicksight_getFlowMetadataCmd)
}
