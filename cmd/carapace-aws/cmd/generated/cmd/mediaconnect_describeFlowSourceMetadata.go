package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeFlowSourceMetadataCmd = &cobra.Command{
	Use:   "describe-flow-source-metadata",
	Short: "The `DescribeFlowSourceMetadata` API is used to view information about the flow's source transport stream and programs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeFlowSourceMetadataCmd).Standalone()

	mediaconnect_describeFlowSourceMetadataCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow.")
	mediaconnect_describeFlowSourceMetadataCmd.MarkFlagRequired("flow-arn")
	mediaconnectCmd.AddCommand(mediaconnect_describeFlowSourceMetadataCmd)
}
