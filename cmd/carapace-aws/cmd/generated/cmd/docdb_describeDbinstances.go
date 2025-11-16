package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbinstancesCmd = &cobra.Command{
	Use:   "describe-dbinstances",
	Short: "Returns information about provisioned Amazon DocumentDB instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbinstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeDbinstancesCmd).Standalone()

		docdb_describeDbinstancesCmd.Flags().String("dbinstance-identifier", "", "The user-provided instance identifier.")
		docdb_describeDbinstancesCmd.Flags().String("filters", "", "A filter that specifies one or more instances to describe.")
		docdb_describeDbinstancesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describeDbinstancesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	docdbCmd.AddCommand(docdb_describeDbinstancesCmd)
}
