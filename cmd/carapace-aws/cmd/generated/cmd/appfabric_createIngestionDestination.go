package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_createIngestionDestinationCmd = &cobra.Command{
	Use:   "create-ingestion-destination",
	Short: "Creates an ingestion destination, which specifies how an application's ingested data is processed by Amazon Web Services AppFabric and where it's delivered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_createIngestionDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_createIngestionDestinationCmd).Standalone()

		appfabric_createIngestionDestinationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_createIngestionDestinationCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appfabric_createIngestionDestinationCmd.Flags().String("destination-configuration", "", "Contains information about the destination of ingested data.")
		appfabric_createIngestionDestinationCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
		appfabric_createIngestionDestinationCmd.Flags().String("processing-configuration", "", "Contains information about how ingested data is processed.")
		appfabric_createIngestionDestinationCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
		appfabric_createIngestionDestinationCmd.MarkFlagRequired("app-bundle-identifier")
		appfabric_createIngestionDestinationCmd.MarkFlagRequired("destination-configuration")
		appfabric_createIngestionDestinationCmd.MarkFlagRequired("ingestion-identifier")
		appfabric_createIngestionDestinationCmd.MarkFlagRequired("processing-configuration")
	})
	appfabricCmd.AddCommand(appfabric_createIngestionDestinationCmd)
}
