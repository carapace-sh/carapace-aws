package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeOrderableDbinstanceOptionsCmd = &cobra.Command{
	Use:   "describe-orderable-dbinstance-options",
	Short: "Returns a list of orderable DB instance options for the specified engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeOrderableDbinstanceOptionsCmd).Standalone()

	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("dbinstance-class", "", "The DB instance class filter value.")
	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("engine", "", "The name of the engine to retrieve DB instance options for.")
	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("engine-version", "", "The engine version filter value.")
	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("license-model", "", "The license model filter value.")
	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions request.")
	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptune_describeOrderableDbinstanceOptionsCmd.Flags().String("vpc", "", "The VPC filter value.")
	neptune_describeOrderableDbinstanceOptionsCmd.MarkFlagRequired("engine")
	neptuneCmd.AddCommand(neptune_describeOrderableDbinstanceOptionsCmd)
}
