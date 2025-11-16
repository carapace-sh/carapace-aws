package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelExportsToTargetCmd = &cobra.Command{
	Use:   "describe-metadata-model-exports-to-target",
	Short: "Returns a paginated list of metadata model exports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelExportsToTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeMetadataModelExportsToTargetCmd).Standalone()

		dms_describeMetadataModelExportsToTargetCmd.Flags().String("filters", "", "Filters applied to the metadata model exports described in the form of key-value pairs.")
		dms_describeMetadataModelExportsToTargetCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
		dms_describeMetadataModelExportsToTargetCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeMetadataModelExportsToTargetCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_describeMetadataModelExportsToTargetCmd.MarkFlagRequired("migration-project-identifier")
	})
	dmsCmd.AddCommand(dms_describeMetadataModelExportsToTargetCmd)
}
