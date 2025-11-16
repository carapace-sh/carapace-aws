package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateIndexCmd = &cobra.Command{
	Use:   "update-index",
	Short: "Updates an existing index in an OpenSearch Serverless collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_updateIndexCmd).Standalone()

		opensearchserverless_updateIndexCmd.Flags().String("id", "", "The unique identifier of the collection containing the index to update.")
		opensearchserverless_updateIndexCmd.Flags().String("index-name", "", "The name of the index to update.")
		opensearchserverless_updateIndexCmd.Flags().String("index-schema", "", "The updated JSON schema definition for the index, including field mappings and settings.")
		opensearchserverless_updateIndexCmd.MarkFlagRequired("id")
		opensearchserverless_updateIndexCmd.MarkFlagRequired("index-name")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateIndexCmd)
}
