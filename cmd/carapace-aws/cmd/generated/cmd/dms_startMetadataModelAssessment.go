package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startMetadataModelAssessmentCmd = &cobra.Command{
	Use:   "start-metadata-model-assessment",
	Short: "Creates a database migration assessment report by assessing the migration complexity for your source database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startMetadataModelAssessmentCmd).Standalone()

	dms_startMetadataModelAssessmentCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
	dms_startMetadataModelAssessmentCmd.Flags().String("selection-rules", "", "A value that specifies the database objects to assess.")
	dms_startMetadataModelAssessmentCmd.MarkFlagRequired("migration-project-identifier")
	dms_startMetadataModelAssessmentCmd.MarkFlagRequired("selection-rules")
	dmsCmd.AddCommand(dms_startMetadataModelAssessmentCmd)
}
