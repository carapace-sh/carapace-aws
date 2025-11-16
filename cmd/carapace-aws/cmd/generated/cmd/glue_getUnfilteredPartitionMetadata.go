package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getUnfilteredPartitionMetadataCmd = &cobra.Command{
	Use:   "get-unfiltered-partition-metadata",
	Short: "Retrieves partition metadata from the Data Catalog that contains unfiltered metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getUnfilteredPartitionMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getUnfilteredPartitionMetadataCmd).Standalone()

		glue_getUnfilteredPartitionMetadataCmd.Flags().String("audit-context", "", "A structure containing Lake Formation audit context information.")
		glue_getUnfilteredPartitionMetadataCmd.Flags().String("catalog-id", "", "The catalog ID where the partition resides.")
		glue_getUnfilteredPartitionMetadataCmd.Flags().String("database-name", "", "(Required) Specifies the name of a database that contains the partition.")
		glue_getUnfilteredPartitionMetadataCmd.Flags().String("partition-values", "", "(Required) A list of partition key values.")
		glue_getUnfilteredPartitionMetadataCmd.Flags().String("query-session-context", "", "A structure used as a protocol between query engines and Lake Formation or Glue.")
		glue_getUnfilteredPartitionMetadataCmd.Flags().String("region", "", "Specified only if the base tables belong to a different Amazon Web Services Region.")
		glue_getUnfilteredPartitionMetadataCmd.Flags().String("supported-permission-types", "", "(Required) A list of supported permission types.")
		glue_getUnfilteredPartitionMetadataCmd.Flags().String("table-name", "", "(Required) Specifies the name of a table that contains the partition.")
		glue_getUnfilteredPartitionMetadataCmd.MarkFlagRequired("catalog-id")
		glue_getUnfilteredPartitionMetadataCmd.MarkFlagRequired("database-name")
		glue_getUnfilteredPartitionMetadataCmd.MarkFlagRequired("partition-values")
		glue_getUnfilteredPartitionMetadataCmd.MarkFlagRequired("supported-permission-types")
		glue_getUnfilteredPartitionMetadataCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_getUnfilteredPartitionMetadataCmd)
}
