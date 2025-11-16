package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelCmd = &cobra.Command{
	Use:   "describe-metadata-model",
	Short: "Gets detailed information about the specified metadata model, including its definition and corresponding converted objects in the target database if applicable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelCmd).Standalone()

	dms_describeMetadataModelCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
	dms_describeMetadataModelCmd.Flags().String("origin", "", "Specifies whether to retrieve metadata from the source or target tree.")
	dms_describeMetadataModelCmd.Flags().String("selection-rules", "", "The JSON string that specifies which metadata model to retrieve.")
	dms_describeMetadataModelCmd.MarkFlagRequired("migration-project-identifier")
	dms_describeMetadataModelCmd.MarkFlagRequired("origin")
	dms_describeMetadataModelCmd.MarkFlagRequired("selection-rules")
	dmsCmd.AddCommand(dms_describeMetadataModelCmd)
}
