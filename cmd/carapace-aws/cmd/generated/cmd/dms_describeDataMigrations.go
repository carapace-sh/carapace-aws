package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeDataMigrationsCmd = &cobra.Command{
	Use:   "describe-data-migrations",
	Short: "Returns information about data migrations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeDataMigrationsCmd).Standalone()

	dms_describeDataMigrationsCmd.Flags().String("filters", "", "Filters applied to the data migrations.")
	dms_describeDataMigrationsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeDataMigrationsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeDataMigrationsCmd.Flags().String("without-settings", "", "An option to set to avoid returning information about settings.")
	dms_describeDataMigrationsCmd.Flags().String("without-statistics", "", "An option to set to avoid returning information about statistics.")
	dmsCmd.AddCommand(dms_describeDataMigrationsCmd)
}
