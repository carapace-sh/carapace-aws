package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelConversionsCmd = &cobra.Command{
	Use:   "describe-metadata-model-conversions",
	Short: "Returns a paginated list of metadata model conversions for a migration project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelConversionsCmd).Standalone()

	dms_describeMetadataModelConversionsCmd.Flags().String("filters", "", "Filters applied to the metadata model conversions described in the form of key-value pairs.")
	dms_describeMetadataModelConversionsCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
	dms_describeMetadataModelConversionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeMetadataModelConversionsCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
	dms_describeMetadataModelConversionsCmd.MarkFlagRequired("migration-project-identifier")
	dmsCmd.AddCommand(dms_describeMetadataModelConversionsCmd)
}
