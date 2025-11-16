package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelAssessmentsCmd = &cobra.Command{
	Use:   "describe-metadata-model-assessments",
	Short: "Returns a paginated list of metadata model assessments for your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelAssessmentsCmd).Standalone()

	dms_describeMetadataModelAssessmentsCmd.Flags().String("filters", "", "Filters applied to the metadata model assessments described in the form of key-value pairs.")
	dms_describeMetadataModelAssessmentsCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
	dms_describeMetadataModelAssessmentsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeMetadataModelAssessmentsCmd.Flags().String("migration-project-identifier", "", "The name or Amazon Resource Name (ARN) of the migration project.")
	dms_describeMetadataModelAssessmentsCmd.MarkFlagRequired("migration-project-identifier")
	dmsCmd.AddCommand(dms_describeMetadataModelAssessmentsCmd)
}
