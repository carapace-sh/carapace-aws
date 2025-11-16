package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbparameterGroupsCmd = &cobra.Command{
	Use:   "describe-dbparameter-groups",
	Short: "Returns a list of `DBParameterGroup` descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbparameterGroupsCmd).Standalone()

	neptune_describeDbparameterGroupsCmd.Flags().String("dbparameter-group-name", "", "The name of a specific DB parameter group to return details for.")
	neptune_describeDbparameterGroupsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	neptune_describeDbparameterGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBParameterGroups` request.")
	neptune_describeDbparameterGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptuneCmd.AddCommand(neptune_describeDbparameterGroupsCmd)
}
