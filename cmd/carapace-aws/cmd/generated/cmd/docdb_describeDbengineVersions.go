package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbengineVersionsCmd = &cobra.Command{
	Use:   "describe-dbengine-versions",
	Short: "Returns a list of the available engines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbengineVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeDbengineVersionsCmd).Standalone()

		docdb_describeDbengineVersionsCmd.Flags().String("dbparameter-group-family", "", "The name of a specific parameter group family to return details for.")
		docdb_describeDbengineVersionsCmd.Flags().Bool("default-only", false, "Indicates that only the default version of the specified engine or engine and major version combination is returned.")
		docdb_describeDbengineVersionsCmd.Flags().String("engine", "", "The database engine to return.")
		docdb_describeDbengineVersionsCmd.Flags().String("engine-version", "", "The database engine version to return.")
		docdb_describeDbengineVersionsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		docdb_describeDbengineVersionsCmd.Flags().String("list-supported-character-sets", "", "If this parameter is specified and the requested engine supports the `CharacterSetName` parameter for `CreateDBInstance`, the response includes a list of supported character sets for each engine version.")
		docdb_describeDbengineVersionsCmd.Flags().String("list-supported-timezones", "", "If this parameter is specified and the requested engine supports the `TimeZone` parameter for `CreateDBInstance`, the response includes a list of supported time zones for each engine version.")
		docdb_describeDbengineVersionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describeDbengineVersionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		docdb_describeDbengineVersionsCmd.Flags().Bool("no-default-only", false, "Indicates that only the default version of the specified engine or engine and major version combination is returned.")
		docdb_describeDbengineVersionsCmd.Flag("no-default-only").Hidden = true
	})
	docdbCmd.AddCommand(docdb_describeDbengineVersionsCmd)
}
