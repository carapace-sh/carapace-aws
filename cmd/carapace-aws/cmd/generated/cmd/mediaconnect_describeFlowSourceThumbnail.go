package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeFlowSourceThumbnailCmd = &cobra.Command{
	Use:   "describe-flow-source-thumbnail",
	Short: "Describes the thumbnail for the flow source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeFlowSourceThumbnailCmd).Standalone()

	mediaconnect_describeFlowSourceThumbnailCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow.")
	mediaconnect_describeFlowSourceThumbnailCmd.MarkFlagRequired("flow-arn")
	mediaconnectCmd.AddCommand(mediaconnect_describeFlowSourceThumbnailCmd)
}
