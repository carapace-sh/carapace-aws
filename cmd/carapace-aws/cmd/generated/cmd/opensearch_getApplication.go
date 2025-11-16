package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Retrieves the configuration and status of an existing OpenSearch application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_getApplicationCmd).Standalone()

		opensearch_getApplicationCmd.Flags().String("id", "", "The unique identifier of the OpenSearch application to retrieve.")
		opensearch_getApplicationCmd.MarkFlagRequired("id")
	})
	opensearchCmd.AddCommand(opensearch_getApplicationCmd)
}
