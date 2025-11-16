package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_getResourceMetadataCmd = &cobra.Command{
	Use:   "get-resource-metadata",
	Short: "Retrieve the metadata for different features.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_getResourceMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_getResourceMetadataCmd).Standalone()

		pi_getResourceMetadataCmd.Flags().String("identifier", "", "An immutable identifier for a data source that is unique for an Amazon Web Services Region.")
		pi_getResourceMetadataCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights returns metrics.")
		pi_getResourceMetadataCmd.MarkFlagRequired("identifier")
		pi_getResourceMetadataCmd.MarkFlagRequired("service-type")
	})
	piCmd.AddCommand(pi_getResourceMetadataCmd)
}
