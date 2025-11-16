package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createMigrationProjectCmd = &cobra.Command{
	Use:   "create-migration-project",
	Short: "Creates the migration project using the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createMigrationProjectCmd).Standalone()

	dms_createMigrationProjectCmd.Flags().String("description", "", "A user-friendly description of the migration project.")
	dms_createMigrationProjectCmd.Flags().String("instance-profile-identifier", "", "The identifier of the associated instance profile.")
	dms_createMigrationProjectCmd.Flags().String("migration-project-name", "", "A user-friendly name for the migration project.")
	dms_createMigrationProjectCmd.Flags().String("schema-conversion-application-attributes", "", "The schema conversion application attributes, including the Amazon S3 bucket name and Amazon S3 role ARN.")
	dms_createMigrationProjectCmd.Flags().String("source-data-provider-descriptors", "", "Information about the source data provider, including the name, ARN, and Secrets Manager parameters.")
	dms_createMigrationProjectCmd.Flags().String("tags", "", "One or more tags to be assigned to the migration project.")
	dms_createMigrationProjectCmd.Flags().String("target-data-provider-descriptors", "", "Information about the target data provider, including the name, ARN, and Amazon Web Services Secrets Manager parameters.")
	dms_createMigrationProjectCmd.Flags().String("transformation-rules", "", "The settings in JSON format for migration rules.")
	dms_createMigrationProjectCmd.MarkFlagRequired("instance-profile-identifier")
	dms_createMigrationProjectCmd.MarkFlagRequired("source-data-provider-descriptors")
	dms_createMigrationProjectCmd.MarkFlagRequired("target-data-provider-descriptors")
	dmsCmd.AddCommand(dms_createMigrationProjectCmd)
}
