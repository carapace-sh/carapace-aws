package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_getIngestionCmd = &cobra.Command{
	Use:   "get-ingestion",
	Short: "Returns information about an ingestion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_getIngestionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_getIngestionCmd).Standalone()

		appfabric_getIngestionCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_getIngestionCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
		appfabric_getIngestionCmd.MarkFlagRequired("app-bundle-identifier")
		appfabric_getIngestionCmd.MarkFlagRequired("ingestion-identifier")
	})
	appfabricCmd.AddCommand(appfabric_getIngestionCmd)
}
