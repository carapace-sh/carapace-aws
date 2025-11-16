package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbengineVersionsCmd = &cobra.Command{
	Use:   "describe-dbengine-versions",
	Short: "Returns a list of the available DB engines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbengineVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeDbengineVersionsCmd).Standalone()

		neptune_describeDbengineVersionsCmd.Flags().String("dbparameter-group-family", "", "The name of a specific DB parameter group family to return details for.")
		neptune_describeDbengineVersionsCmd.Flags().Bool("default-only", false, "Indicates that only the default version of the specified engine or engine and major version combination is returned.")
		neptune_describeDbengineVersionsCmd.Flags().String("engine", "", "The database engine to return.")
		neptune_describeDbengineVersionsCmd.Flags().String("engine-version", "", "The database engine version to return.")
		neptune_describeDbengineVersionsCmd.Flags().String("filters", "", "Not currently supported.")
		neptune_describeDbengineVersionsCmd.Flags().String("list-supported-character-sets", "", "If this parameter is specified and the requested engine supports the `CharacterSetName` parameter for `CreateDBInstance`, the response includes a list of supported character sets for each engine version.")
		neptune_describeDbengineVersionsCmd.Flags().String("list-supported-timezones", "", "If this parameter is specified and the requested engine supports the `TimeZone` parameter for `CreateDBInstance`, the response includes a list of supported time zones for each engine version.")
		neptune_describeDbengineVersionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		neptune_describeDbengineVersionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		neptune_describeDbengineVersionsCmd.Flags().Bool("no-default-only", false, "Indicates that only the default version of the specified engine or engine and major version combination is returned.")
		neptune_describeDbengineVersionsCmd.Flag("no-default-only").Hidden = true
	})
	neptuneCmd.AddCommand(neptune_describeDbengineVersionsCmd)
}
