package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyMigrationProjectCmd = &cobra.Command{
	Use:   "modify-migration-project",
	Short: "Modifies the specified migration project using the provided parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyMigrationProjectCmd).Standalone()

	dms_modifyMigrationProjectCmd.Flags().String("description", "", "A user-friendly description of the migration project.")
	dms_modifyMigrationProjectCmd.Flags().String("instance-profile-identifier", "", "The name or Amazon Resource Name (ARN) for the instance profile.")
	dms_modifyMigrationProjectCmd.Flags().String("migration-project-identifier", "", "The identifier of the migration project.")
	dms_modifyMigrationProjectCmd.Flags().String("migration-project-name", "", "A user-friendly name for the migration project.")
	dms_modifyMigrationProjectCmd.Flags().String("schema-conversion-application-attributes", "", "The schema conversion application attributes, including the Amazon S3 bucket name and Amazon S3 role ARN.")
	dms_modifyMigrationProjectCmd.Flags().String("source-data-provider-descriptors", "", "Information about the source data provider, including the name, ARN, and Amazon Web Services Secrets Manager parameters.")
	dms_modifyMigrationProjectCmd.Flags().String("target-data-provider-descriptors", "", "Information about the target data provider, including the name, ARN, and Amazon Web Services Secrets Manager parameters.")
	dms_modifyMigrationProjectCmd.Flags().String("transformation-rules", "", "The settings in JSON format for migration rules.")
	dms_modifyMigrationProjectCmd.MarkFlagRequired("migration-project-identifier")
	dmsCmd.AddCommand(dms_modifyMigrationProjectCmd)
}
