package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbengineVersionsCmd = &cobra.Command{
	Use:   "describe-dbengine-versions",
	Short: "Describes the properties of specific versions of DB engines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbengineVersionsCmd).Standalone()

	rds_describeDbengineVersionsCmd.Flags().String("dbparameter-group-family", "", "The name of a specific DB parameter group family to return details for.")
	rds_describeDbengineVersionsCmd.Flags().Bool("default-only", false, "Specifies whether to return only the default version of the specified engine or the engine and major version combination.")
	rds_describeDbengineVersionsCmd.Flags().String("engine", "", "The database engine to return version details for.")
	rds_describeDbengineVersionsCmd.Flags().String("engine-version", "", "A specific database engine version to return details for.")
	rds_describeDbengineVersionsCmd.Flags().String("filters", "", "A filter that specifies one or more DB engine versions to describe.")
	rds_describeDbengineVersionsCmd.Flags().String("include-all", "", "Specifies whether to also list the engine versions that aren't available.")
	rds_describeDbengineVersionsCmd.Flags().String("list-supported-character-sets", "", "Specifies whether to list the supported character sets for each engine version.")
	rds_describeDbengineVersionsCmd.Flags().String("list-supported-timezones", "", "Specifies whether to list the supported time zones for each engine version.")
	rds_describeDbengineVersionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	rds_describeDbengineVersionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeDbengineVersionsCmd.Flags().Bool("no-default-only", false, "Specifies whether to return only the default version of the specified engine or the engine and major version combination.")
	rds_describeDbengineVersionsCmd.Flag("no-default-only").Hidden = true
	rdsCmd.AddCommand(rds_describeDbengineVersionsCmd)
}
