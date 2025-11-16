package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbmajorEngineVersionsCmd = &cobra.Command{
	Use:   "describe-dbmajor-engine-versions",
	Short: "Describes the properties of specific major versions of DB engines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbmajorEngineVersionsCmd).Standalone()

	rds_describeDbmajorEngineVersionsCmd.Flags().String("engine", "", "The database engine to return major version details for.")
	rds_describeDbmajorEngineVersionsCmd.Flags().String("major-engine-version", "", "A specific database major engine version to return details for.")
	rds_describeDbmajorEngineVersionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	rds_describeDbmajorEngineVersionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rdsCmd.AddCommand(rds_describeDbmajorEngineVersionsCmd)
}
