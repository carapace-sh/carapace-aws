package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_getIngestionDestinationCmd = &cobra.Command{
	Use:   "get-ingestion-destination",
	Short: "Returns information about an ingestion destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_getIngestionDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_getIngestionDestinationCmd).Standalone()

		appfabric_getIngestionDestinationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_getIngestionDestinationCmd.Flags().String("ingestion-destination-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion destination to use for the request.")
		appfabric_getIngestionDestinationCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
		appfabric_getIngestionDestinationCmd.MarkFlagRequired("app-bundle-identifier")
		appfabric_getIngestionDestinationCmd.MarkFlagRequired("ingestion-destination-identifier")
		appfabric_getIngestionDestinationCmd.MarkFlagRequired("ingestion-identifier")
	})
	appfabricCmd.AddCommand(appfabric_getIngestionDestinationCmd)
}
