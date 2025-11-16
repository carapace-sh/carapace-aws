package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_deleteIngestionCmd = &cobra.Command{
	Use:   "delete-ingestion",
	Short: "Deletes an ingestion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_deleteIngestionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_deleteIngestionCmd).Standalone()

		appfabric_deleteIngestionCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_deleteIngestionCmd.Flags().String("ingestion-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to use for the request.")
		appfabric_deleteIngestionCmd.MarkFlagRequired("app-bundle-identifier")
		appfabric_deleteIngestionCmd.MarkFlagRequired("ingestion-identifier")
	})
	appfabricCmd.AddCommand(appfabric_deleteIngestionCmd)
}
