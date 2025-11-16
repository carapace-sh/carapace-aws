package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteMigrationProjectCmd = &cobra.Command{
	Use:   "delete-migration-project",
	Short: "Deletes the specified migration project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteMigrationProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteMigrationProjectCmd).Standalone()

		dms_deleteMigrationProjectCmd.Flags().String("migration-project-identifier", "", "The name or Amazon Resource Name (ARN) of the migration project to delete.")
		dms_deleteMigrationProjectCmd.MarkFlagRequired("migration-project-identifier")
	})
	dmsCmd.AddCommand(dms_deleteMigrationProjectCmd)
}
