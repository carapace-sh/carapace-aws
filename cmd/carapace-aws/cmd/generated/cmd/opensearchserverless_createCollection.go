package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_createCollectionCmd = &cobra.Command{
	Use:   "create-collection",
	Short: "Creates a new OpenSearch Serverless collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_createCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_createCollectionCmd).Standalone()

		opensearchserverless_createCollectionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearchserverless_createCollectionCmd.Flags().String("description", "", "Description of the collection.")
		opensearchserverless_createCollectionCmd.Flags().String("name", "", "Name of the collection.")
		opensearchserverless_createCollectionCmd.Flags().String("standby-replicas", "", "Indicates whether standby replicas should be used for a collection.")
		opensearchserverless_createCollectionCmd.Flags().String("tags", "", "An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Serverless collection.")
		opensearchserverless_createCollectionCmd.Flags().String("type", "", "The type of collection.")
		opensearchserverless_createCollectionCmd.MarkFlagRequired("name")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_createCollectionCmd)
}
