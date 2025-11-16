package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbparametersCmd = &cobra.Command{
	Use:   "describe-dbparameters",
	Short: "Returns the detailed parameter list for a particular DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbparametersCmd).Standalone()

	rds_describeDbparametersCmd.Flags().String("dbparameter-group-name", "", "The name of a specific DB parameter group to return details for.")
	rds_describeDbparametersCmd.Flags().String("filters", "", "A filter that specifies one or more DB parameters to describe.")
	rds_describeDbparametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBParameters` request.")
	rds_describeDbparametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeDbparametersCmd.Flags().String("source", "", "The parameter types to return.")
	rds_describeDbparametersCmd.MarkFlagRequired("dbparameter-group-name")
	rdsCmd.AddCommand(rds_describeDbparametersCmd)
}
