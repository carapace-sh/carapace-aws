package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getUnfilteredPartitionsMetadataCmd = &cobra.Command{
	Use:   "get-unfiltered-partitions-metadata",
	Short: "Retrieves partition metadata from the Data Catalog that contains unfiltered metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getUnfilteredPartitionsMetadataCmd).Standalone()

	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("audit-context", "", "A structure containing Lake Formation audit context information.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("expression", "", "An expression that filters the partitions to be returned.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("max-results", "", "The maximum number of partitions to return in a single response.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve these partitions.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("query-session-context", "", "A structure used as a protocol between query engines and Lake Formation or Glue.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("region", "", "Specified only if the base tables belong to a different Amazon Web Services Region.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("segment", "", "The segment of the table's partitions to scan in this request.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("supported-permission-types", "", "A list of supported permission types.")
	glue_getUnfilteredPartitionsMetadataCmd.Flags().String("table-name", "", "The name of the table that contains the partition.")
	glue_getUnfilteredPartitionsMetadataCmd.MarkFlagRequired("catalog-id")
	glue_getUnfilteredPartitionsMetadataCmd.MarkFlagRequired("database-name")
	glue_getUnfilteredPartitionsMetadataCmd.MarkFlagRequired("supported-permission-types")
	glue_getUnfilteredPartitionsMetadataCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_getUnfilteredPartitionsMetadataCmd)
}
