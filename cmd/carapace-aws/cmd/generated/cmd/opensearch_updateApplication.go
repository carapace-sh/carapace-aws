package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates the configuration and settings of an existing OpenSearch application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updateApplicationCmd).Standalone()

	opensearch_updateApplicationCmd.Flags().String("app-configs", "", "The configuration settings to modify for the OpenSearch application.")
	opensearch_updateApplicationCmd.Flags().String("data-sources", "", "The data sources to associate with the OpenSearch application.")
	opensearch_updateApplicationCmd.Flags().String("id", "", "The unique identifier for the OpenSearch application to be updated.")
	opensearch_updateApplicationCmd.MarkFlagRequired("id")
	opensearchCmd.AddCommand(opensearch_updateApplicationCmd)
}
