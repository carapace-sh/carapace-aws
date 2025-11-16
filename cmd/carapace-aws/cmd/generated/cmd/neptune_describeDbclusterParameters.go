package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbclusterParametersCmd = &cobra.Command{
	Use:   "describe-dbcluster-parameters",
	Short: "Returns the detailed parameter list for a particular DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbclusterParametersCmd).Standalone()

	neptune_describeDbclusterParametersCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of a specific DB cluster parameter group to return parameter details for.")
	neptune_describeDbclusterParametersCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	neptune_describeDbclusterParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterParameters` request.")
	neptune_describeDbclusterParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptune_describeDbclusterParametersCmd.Flags().String("source", "", "A value that indicates to return only parameters for a specific source.")
	neptune_describeDbclusterParametersCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	neptuneCmd.AddCommand(neptune_describeDbclusterParametersCmd)
}
