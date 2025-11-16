package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateViewMetadataCmd = &cobra.Command{
	Use:   "update-view-metadata",
	Short: "Updates the view metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateViewMetadataCmd).Standalone()

	connect_updateViewMetadataCmd.Flags().String("description", "", "The description of the view.")
	connect_updateViewMetadataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateViewMetadataCmd.Flags().String("name", "", "The name of the view.")
	connect_updateViewMetadataCmd.Flags().String("view-id", "", "The identifier of the view.")
	connect_updateViewMetadataCmd.MarkFlagRequired("instance-id")
	connect_updateViewMetadataCmd.MarkFlagRequired("view-id")
	connectCmd.AddCommand(connect_updateViewMetadataCmd)
}
