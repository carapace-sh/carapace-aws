package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_updateIngestionDestinationCmd = &cobra.Command{
	Use:   "update-ingestion-destination",
	Short: "Updates an ingestion destination, which specifies how an application's ingested data is processed by Amazon Web Services AppFabric and where it's delivered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_updateIngestionDestinationCmd).Standalone()

	appfabric_updateIngestionDestinationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_updateIngestionDestinationCmd.Flags().String("destination-configuration", "", "Contains information about the destination of ingested data.")
	appfabric_updateIngestionDestinationCmd.Flags().String("ingestion-destination-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion destination to use for the request.")
	appfabric_updateIngestionDestinationCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
	appfabric_updateIngestionDestinationCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_updateIngestionDestinationCmd.MarkFlagRequired("destination-configuration")
	appfabric_updateIngestionDestinationCmd.MarkFlagRequired("ingestion-destination-identifier")
	appfabric_updateIngestionDestinationCmd.MarkFlagRequired("ingestion-identifier")
	appfabricCmd.AddCommand(appfabric_updateIngestionDestinationCmd)
}
