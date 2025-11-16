package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getEntityRecordsCmd = &cobra.Command{
	Use:   "get-entity-records",
	Short: "This API is used to query preview data from a given connection type or from a native Amazon S3 based Glue Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getEntityRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getEntityRecordsCmd).Standalone()

		glue_getEntityRecordsCmd.Flags().String("catalog-id", "", "The catalog ID of the catalog that contains the connection.")
		glue_getEntityRecordsCmd.Flags().String("connection-name", "", "The name of the connection that contains the connection type credentials.")
		glue_getEntityRecordsCmd.Flags().String("connection-options", "", "Connector options that are required to query the data.")
		glue_getEntityRecordsCmd.Flags().String("data-store-api-version", "", "The API version of the SaaS connector.")
		glue_getEntityRecordsCmd.Flags().String("entity-name", "", "Name of the entity that we want to query the preview data from the given connection type.")
		glue_getEntityRecordsCmd.Flags().String("filter-predicate", "", "A filter predicate that you can apply in the query request.")
		glue_getEntityRecordsCmd.Flags().String("limit", "", "Limits the number of records fetched with the request.")
		glue_getEntityRecordsCmd.Flags().String("next-token", "", "A continuation token, included if this is a continuation call.")
		glue_getEntityRecordsCmd.Flags().String("order-by", "", "A parameter that orders the response preview data.")
		glue_getEntityRecordsCmd.Flags().String("selected-fields", "", "List of fields that we want to fetch as part of preview data.")
		glue_getEntityRecordsCmd.MarkFlagRequired("entity-name")
		glue_getEntityRecordsCmd.MarkFlagRequired("limit")
	})
	glueCmd.AddCommand(glue_getEntityRecordsCmd)
}
