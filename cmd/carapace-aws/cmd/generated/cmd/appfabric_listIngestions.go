package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_listIngestionsCmd = &cobra.Command{
	Use:   "list-ingestions",
	Short: "Returns a list of all ingestions configured for an app bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_listIngestionsCmd).Standalone()

	appfabric_listIngestionsCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_listIngestionsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	appfabric_listIngestionsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	appfabric_listIngestionsCmd.MarkFlagRequired("app-bundle-identifier")
	appfabricCmd.AddCommand(appfabric_listIngestionsCmd)
}
