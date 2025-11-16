package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_deleteIndexCmd = &cobra.Command{
	Use:   "delete-index",
	Short: "Deletes an index from an OpenSearch Serverless collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_deleteIndexCmd).Standalone()

	opensearchserverless_deleteIndexCmd.Flags().String("id", "", "The unique identifier of the collection containing the index to delete.")
	opensearchserverless_deleteIndexCmd.Flags().String("index-name", "", "The name of the index to delete.")
	opensearchserverless_deleteIndexCmd.MarkFlagRequired("id")
	opensearchserverless_deleteIndexCmd.MarkFlagRequired("index-name")
	opensearchserverlessCmd.AddCommand(opensearchserverless_deleteIndexCmd)
}
