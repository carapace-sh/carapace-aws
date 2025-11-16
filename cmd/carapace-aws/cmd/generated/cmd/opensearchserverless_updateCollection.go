package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateCollectionCmd = &cobra.Command{
	Use:   "update-collection",
	Short: "Updates an OpenSearch Serverless collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateCollectionCmd).Standalone()

	opensearchserverless_updateCollectionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_updateCollectionCmd.Flags().String("description", "", "A description of the collection.")
	opensearchserverless_updateCollectionCmd.Flags().String("id", "", "The unique identifier of the collection.")
	opensearchserverless_updateCollectionCmd.MarkFlagRequired("id")
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateCollectionCmd)
}
