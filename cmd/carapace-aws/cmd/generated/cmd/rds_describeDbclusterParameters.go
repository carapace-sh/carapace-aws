package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclusterParametersCmd = &cobra.Command{
	Use:   "describe-dbcluster-parameters",
	Short: "Returns the detailed parameter list for a particular DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclusterParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbclusterParametersCmd).Standalone()

		rds_describeDbclusterParametersCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of a specific DB cluster parameter group to return parameter details for.")
		rds_describeDbclusterParametersCmd.Flags().String("filters", "", "A filter that specifies one or more DB cluster parameters to describe.")
		rds_describeDbclusterParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterParameters` request.")
		rds_describeDbclusterParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeDbclusterParametersCmd.Flags().String("source", "", "A specific source to return parameters for.")
		rds_describeDbclusterParametersCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	})
	rdsCmd.AddCommand(rds_describeDbclusterParametersCmd)
}
