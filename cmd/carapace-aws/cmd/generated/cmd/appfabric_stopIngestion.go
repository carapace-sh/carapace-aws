package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_stopIngestionCmd = &cobra.Command{
	Use:   "stop-ingestion",
	Short: "Stops (disables) an ingestion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_stopIngestionCmd).Standalone()

	appfabric_stopIngestionCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_stopIngestionCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
	appfabric_stopIngestionCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_stopIngestionCmd.MarkFlagRequired("ingestion-identifier")
	appfabricCmd.AddCommand(appfabric_stopIngestionCmd)
}
