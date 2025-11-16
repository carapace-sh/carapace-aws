package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes a specified OpenSearch application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deleteApplicationCmd).Standalone()

	opensearch_deleteApplicationCmd.Flags().String("id", "", "The unique identifier of the OpenSearch application to delete.")
	opensearch_deleteApplicationCmd.MarkFlagRequired("id")
	opensearchCmd.AddCommand(opensearch_deleteApplicationCmd)
}
