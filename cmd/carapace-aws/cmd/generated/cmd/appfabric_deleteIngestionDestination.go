package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_deleteIngestionDestinationCmd = &cobra.Command{
	Use:   "delete-ingestion-destination",
	Short: "Deletes an ingestion destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_deleteIngestionDestinationCmd).Standalone()

	appfabric_deleteIngestionDestinationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_deleteIngestionDestinationCmd.Flags().String("ingestion-destination-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion destination to use for the request.")
	appfabric_deleteIngestionDestinationCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
	appfabric_deleteIngestionDestinationCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_deleteIngestionDestinationCmd.MarkFlagRequired("ingestion-destination-identifier")
	appfabric_deleteIngestionDestinationCmd.MarkFlagRequired("ingestion-identifier")
	appfabricCmd.AddCommand(appfabric_deleteIngestionDestinationCmd)
}
