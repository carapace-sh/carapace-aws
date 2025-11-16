package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getUnfilteredTableMetadataCmd = &cobra.Command{
	Use:   "get-unfiltered-table-metadata",
	Short: "Allows a third-party analytical engine to retrieve unfiltered table metadata from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getUnfilteredTableMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getUnfilteredTableMetadataCmd).Standalone()

		glue_getUnfilteredTableMetadataCmd.Flags().String("audit-context", "", "A structure containing Lake Formation audit context information.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("catalog-id", "", "The catalog ID where the table resides.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("database-name", "", "(Required) Specifies the name of a database that contains the table.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("name", "", "(Required) Specifies the name of a table for which you are requesting metadata.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("parent-resource-arn", "", "The resource ARN of the view.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("permissions", "", "The Lake Formation data permissions of the caller on the table.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("query-session-context", "", "A structure used as a protocol between query engines and Lake Formation or Glue.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("region", "", "Specified only if the base tables belong to a different Amazon Web Services Region.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("root-resource-arn", "", "The resource ARN of the root view in a chain of nested views.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("supported-dialect", "", "A structure specifying the dialect and dialect version used by the query engine.")
		glue_getUnfilteredTableMetadataCmd.Flags().String("supported-permission-types", "", "Indicates the level of filtering a third-party analytical engine is capable of enforcing when calling the `GetUnfilteredTableMetadata` API operation.")
		glue_getUnfilteredTableMetadataCmd.MarkFlagRequired("catalog-id")
		glue_getUnfilteredTableMetadataCmd.MarkFlagRequired("database-name")
		glue_getUnfilteredTableMetadataCmd.MarkFlagRequired("name")
		glue_getUnfilteredTableMetadataCmd.MarkFlagRequired("supported-permission-types")
	})
	glueCmd.AddCommand(glue_getUnfilteredTableMetadataCmd)
}
