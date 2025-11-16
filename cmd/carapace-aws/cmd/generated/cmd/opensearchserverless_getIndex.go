package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_getIndexCmd = &cobra.Command{
	Use:   "get-index",
	Short: "Retrieves information about an index in an OpenSearch Serverless collection, including its schema definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_getIndexCmd).Standalone()

	opensearchserverless_getIndexCmd.Flags().String("id", "", "The unique identifier of the collection containing the index.")
	opensearchserverless_getIndexCmd.Flags().String("index-name", "", "The name of the index to retrieve information about.")
	opensearchserverless_getIndexCmd.MarkFlagRequired("id")
	opensearchserverless_getIndexCmd.MarkFlagRequired("index-name")
	opensearchserverlessCmd.AddCommand(opensearchserverless_getIndexCmd)
}
