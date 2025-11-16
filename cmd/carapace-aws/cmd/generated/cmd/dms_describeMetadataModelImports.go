package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelImportsCmd = &cobra.Command{
	Use:   "describe-metadata-model-imports",
	Short: "Returns a paginated list of metadata model imports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelImportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeMetadataModelImportsCmd).Standalone()

		dms_describeMetadataModelImportsCmd.Flags().String("filters", "", "Filters applied to the metadata model imports described in the form of key-value pairs.")
		dms_describeMetadataModelImportsCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
		dms_describeMetadataModelImportsCmd.Flags().String("max-records", "", "A paginated list of metadata model imports.")
		dms_describeMetadataModelImportsCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_describeMetadataModelImportsCmd.MarkFlagRequired("migration-project-identifier")
	})
	dmsCmd.AddCommand(dms_describeMetadataModelImportsCmd)
}
