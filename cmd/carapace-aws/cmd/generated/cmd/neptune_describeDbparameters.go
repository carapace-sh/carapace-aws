package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbparametersCmd = &cobra.Command{
	Use:   "describe-dbparameters",
	Short: "Returns the detailed parameter list for a particular DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbparametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeDbparametersCmd).Standalone()

		neptune_describeDbparametersCmd.Flags().String("dbparameter-group-name", "", "The name of a specific DB parameter group to return details for.")
		neptune_describeDbparametersCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		neptune_describeDbparametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBParameters` request.")
		neptune_describeDbparametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		neptune_describeDbparametersCmd.Flags().String("source", "", "The parameter types to return.")
		neptune_describeDbparametersCmd.MarkFlagRequired("dbparameter-group-name")
	})
	neptuneCmd.AddCommand(neptune_describeDbparametersCmd)
}
