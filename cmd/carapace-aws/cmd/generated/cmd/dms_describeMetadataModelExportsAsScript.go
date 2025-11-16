package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelExportsAsScriptCmd = &cobra.Command{
	Use:   "describe-metadata-model-exports-as-script",
	Short: "Returns a paginated list of metadata model exports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelExportsAsScriptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeMetadataModelExportsAsScriptCmd).Standalone()

		dms_describeMetadataModelExportsAsScriptCmd.Flags().String("filters", "", "Filters applied to the metadata model exports described in the form of key-value pairs.")
		dms_describeMetadataModelExportsAsScriptCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
		dms_describeMetadataModelExportsAsScriptCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeMetadataModelExportsAsScriptCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_describeMetadataModelExportsAsScriptCmd.MarkFlagRequired("migration-project-identifier")
	})
	dmsCmd.AddCommand(dms_describeMetadataModelExportsAsScriptCmd)
}
