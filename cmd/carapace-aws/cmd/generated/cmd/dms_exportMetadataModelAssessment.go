package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_exportMetadataModelAssessmentCmd = &cobra.Command{
	Use:   "export-metadata-model-assessment",
	Short: "Saves a copy of a database migration assessment report to your Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_exportMetadataModelAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_exportMetadataModelAssessmentCmd).Standalone()

		dms_exportMetadataModelAssessmentCmd.Flags().String("assessment-report-types", "", "The file format of the assessment file.")
		dms_exportMetadataModelAssessmentCmd.Flags().String("file-name", "", "The name of the assessment file to create in your Amazon S3 bucket.")
		dms_exportMetadataModelAssessmentCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_exportMetadataModelAssessmentCmd.Flags().String("selection-rules", "", "A value that specifies the database objects to assess.")
		dms_exportMetadataModelAssessmentCmd.MarkFlagRequired("migration-project-identifier")
		dms_exportMetadataModelAssessmentCmd.MarkFlagRequired("selection-rules")
	})
	dmsCmd.AddCommand(dms_exportMetadataModelAssessmentCmd)
}
