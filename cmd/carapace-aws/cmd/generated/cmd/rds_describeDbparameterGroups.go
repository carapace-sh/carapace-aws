package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbparameterGroupsCmd = &cobra.Command{
	Use:   "describe-dbparameter-groups",
	Short: "Returns a list of `DBParameterGroup` descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbparameterGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbparameterGroupsCmd).Standalone()

		rds_describeDbparameterGroupsCmd.Flags().String("dbparameter-group-name", "", "The name of a specific DB parameter group to return details for.")
		rds_describeDbparameterGroupsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeDbparameterGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBParameterGroups` request.")
		rds_describeDbparameterGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbparameterGroupsCmd)
}
