package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an OpenSearch UI application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_createApplicationCmd).Standalone()

		opensearch_createApplicationCmd.Flags().String("app-configs", "", "Configuration settings for the OpenSearch application, including administrative options.")
		opensearch_createApplicationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearch_createApplicationCmd.Flags().String("data-sources", "", "The data sources to link to the OpenSearch application.")
		opensearch_createApplicationCmd.Flags().String("iam-identity-center-options", "", "Configuration settings for integrating Amazon Web Services IAM Identity Center with the OpenSearch application.")
		opensearch_createApplicationCmd.Flags().String("name", "", "The unique name of the OpenSearch application.")
		opensearch_createApplicationCmd.Flags().String("tag-list", "", "")
		opensearch_createApplicationCmd.MarkFlagRequired("name")
	})
	opensearchCmd.AddCommand(opensearch_createApplicationCmd)
}
