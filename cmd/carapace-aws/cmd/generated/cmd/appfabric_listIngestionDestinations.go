package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_listIngestionDestinationsCmd = &cobra.Command{
	Use:   "list-ingestion-destinations",
	Short: "Returns a list of all ingestion destinations configured for an ingestion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_listIngestionDestinationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_listIngestionDestinationsCmd).Standalone()

		appfabric_listIngestionDestinationsCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_listIngestionDestinationsCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
		appfabric_listIngestionDestinationsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		appfabric_listIngestionDestinationsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
		appfabric_listIngestionDestinationsCmd.MarkFlagRequired("app-bundle-identifier")
		appfabric_listIngestionDestinationsCmd.MarkFlagRequired("ingestion-identifier")
	})
	appfabricCmd.AddCommand(appfabric_listIngestionDestinationsCmd)
}
