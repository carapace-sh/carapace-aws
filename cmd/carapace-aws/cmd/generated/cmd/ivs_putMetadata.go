package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_putMetadataCmd = &cobra.Command{
	Use:   "put-metadata",
	Short: "Inserts metadata into the active stream of the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_putMetadataCmd).Standalone()

	ivs_putMetadataCmd.Flags().String("channel-arn", "", "ARN of the channel into which metadata is inserted.")
	ivs_putMetadataCmd.Flags().String("metadata", "", "Metadata to insert into the stream.")
	ivs_putMetadataCmd.MarkFlagRequired("channel-arn")
	ivs_putMetadataCmd.MarkFlagRequired("metadata")
	ivsCmd.AddCommand(ivs_putMetadataCmd)
}
