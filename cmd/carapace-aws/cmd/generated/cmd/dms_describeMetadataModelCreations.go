package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMetadataModelCreationsCmd = &cobra.Command{
	Use:   "describe-metadata-model-creations",
	Short: "Returns a paginated list of metadata model creation requests for a migration project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMetadataModelCreationsCmd).Standalone()

	dms_describeMetadataModelCreationsCmd.Flags().String("filters", "", "Filters applied to the metadata model creation requests described in the form of key-value pairs.")
	dms_describeMetadataModelCreationsCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of metadata model creation requests.")
	dms_describeMetadataModelCreationsCmd.Flags().String("max-records", "", "The maximum number of metadata model creation requests to include in the response.")
	dms_describeMetadataModelCreationsCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
	dms_describeMetadataModelCreationsCmd.MarkFlagRequired("migration-project-identifier")
	dmsCmd.AddCommand(dms_describeMetadataModelCreationsCmd)
}
