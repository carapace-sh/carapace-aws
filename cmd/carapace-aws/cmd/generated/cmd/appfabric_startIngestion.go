package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_startIngestionCmd = &cobra.Command{
	Use:   "start-ingestion",
	Short: "Starts (enables) an ingestion, which collects data from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_startIngestionCmd).Standalone()

	appfabric_startIngestionCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_startIngestionCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
	appfabric_startIngestionCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_startIngestionCmd.MarkFlagRequired("ingestion-identifier")
	appfabricCmd.AddCommand(appfabric_startIngestionCmd)
}
