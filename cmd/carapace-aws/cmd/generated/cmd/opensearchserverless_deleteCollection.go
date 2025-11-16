package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_deleteCollectionCmd = &cobra.Command{
	Use:   "delete-collection",
	Short: "Deletes an OpenSearch Serverless collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_deleteCollectionCmd).Standalone()

	opensearchserverless_deleteCollectionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_deleteCollectionCmd.Flags().String("id", "", "The unique identifier of the collection.")
	opensearchserverless_deleteCollectionCmd.MarkFlagRequired("id")
	opensearchserverlessCmd.AddCommand(opensearchserverless_deleteCollectionCmd)
}
