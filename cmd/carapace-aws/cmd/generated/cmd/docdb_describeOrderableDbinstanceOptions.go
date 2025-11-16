package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeOrderableDbinstanceOptionsCmd = &cobra.Command{
	Use:   "describe-orderable-dbinstance-options",
	Short: "Returns a list of orderable instance options for the specified engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeOrderableDbinstanceOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeOrderableDbinstanceOptionsCmd).Standalone()

		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("dbinstance-class", "", "The instance class filter value.")
		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("engine", "", "The name of the engine to retrieve instance options for.")
		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("engine-version", "", "The engine version filter value.")
		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("license-model", "", "The license model filter value.")
		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		docdb_describeOrderableDbinstanceOptionsCmd.Flags().String("vpc", "", "The virtual private cloud (VPC) filter value.")
		docdb_describeOrderableDbinstanceOptionsCmd.MarkFlagRequired("engine")
	})
	docdbCmd.AddCommand(docdb_describeOrderableDbinstanceOptionsCmd)
}
