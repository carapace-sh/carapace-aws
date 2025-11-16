package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeMigrationProjectsCmd = &cobra.Command{
	Use:   "describe-migration-projects",
	Short: "Returns a paginated list of migration projects for your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeMigrationProjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeMigrationProjectsCmd).Standalone()

		dms_describeMigrationProjectsCmd.Flags().String("filters", "", "Filters applied to the migration projects described in the form of key-value pairs.")
		dms_describeMigrationProjectsCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
		dms_describeMigrationProjectsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeMigrationProjectsCmd)
}
