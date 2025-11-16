package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbinstancesCmd = &cobra.Command{
	Use:   "describe-dbinstances",
	Short: "Returns information about provisioned instances, and supports pagination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbinstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeDbinstancesCmd).Standalone()

		neptune_describeDbinstancesCmd.Flags().String("dbinstance-identifier", "", "The user-supplied instance identifier.")
		neptune_describeDbinstancesCmd.Flags().String("filters", "", "A filter that specifies one or more DB instances to describe.")
		neptune_describeDbinstancesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBInstances` request.")
		neptune_describeDbinstancesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	neptuneCmd.AddCommand(neptune_describeDbinstancesCmd)
}
