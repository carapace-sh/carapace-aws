package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_createIngestionCmd = &cobra.Command{
	Use:   "create-ingestion",
	Short: "Creates a data ingestion for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_createIngestionCmd).Standalone()

	appfabric_createIngestionCmd.Flags().String("app", "", "The name of the application.")
	appfabric_createIngestionCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_createIngestionCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appfabric_createIngestionCmd.Flags().String("ingestion-type", "", "The ingestion type.")
	appfabric_createIngestionCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	appfabric_createIngestionCmd.Flags().String("tenant-id", "", "The ID of the application tenant.")
	appfabric_createIngestionCmd.MarkFlagRequired("app")
	appfabric_createIngestionCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_createIngestionCmd.MarkFlagRequired("ingestion-type")
	appfabric_createIngestionCmd.MarkFlagRequired("tenant-id")
	appfabricCmd.AddCommand(appfabric_createIngestionCmd)
}
