package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Creates an index within an OpenSearch Serverless collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_createIndexCmd).Standalone()

	opensearchserverless_createIndexCmd.Flags().String("id", "", "The unique identifier of the collection in which to create the index.")
	opensearchserverless_createIndexCmd.Flags().String("index-name", "", "The name of the index to create.")
	opensearchserverless_createIndexCmd.Flags().String("index-schema", "", "The JSON schema definition for the index, including field mappings and settings.")
	opensearchserverless_createIndexCmd.MarkFlagRequired("id")
	opensearchserverless_createIndexCmd.MarkFlagRequired("index-name")
	opensearchserverlessCmd.AddCommand(opensearchserverless_createIndexCmd)
}
