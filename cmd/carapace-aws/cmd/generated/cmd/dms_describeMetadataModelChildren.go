package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelChildrenCmd = &cobra.Command{
	Use:   "describe-metadata-model-children",
	Short: "Gets a list of child metadata models for the specified metadata model in the database hierarchy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelChildrenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeMetadataModelChildrenCmd).Standalone()

		dms_describeMetadataModelChildrenCmd.Flags().String("marker", "", "Specifies the unique pagination token that indicates where the next page should start.")
		dms_describeMetadataModelChildrenCmd.Flags().String("max-records", "", "The maximum number of metadata model children to include in the response.")
		dms_describeMetadataModelChildrenCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_describeMetadataModelChildrenCmd.Flags().String("origin", "", "Specifies whether to retrieve metadata from the source or target tree.")
		dms_describeMetadataModelChildrenCmd.Flags().String("selection-rules", "", "The JSON string that specifies which metadata model's children to retrieve.")
		dms_describeMetadataModelChildrenCmd.MarkFlagRequired("migration-project-identifier")
		dms_describeMetadataModelChildrenCmd.MarkFlagRequired("origin")
		dms_describeMetadataModelChildrenCmd.MarkFlagRequired("selection-rules")
	})
	dmsCmd.AddCommand(dms_describeMetadataModelChildrenCmd)
}
