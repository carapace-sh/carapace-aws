package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbinstancesCmd = &cobra.Command{
	Use:   "describe-dbinstances",
	Short: "Describes provisioned RDS instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbinstancesCmd).Standalone()

	rds_describeDbinstancesCmd.Flags().String("dbinstance-identifier", "", "The user-supplied instance identifier or the Amazon Resource Name (ARN) of the DB instance.")
	rds_describeDbinstancesCmd.Flags().String("filters", "", "A filter that specifies one or more DB instances to describe.")
	rds_describeDbinstancesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBInstances` request.")
	rds_describeDbinstancesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rdsCmd.AddCommand(rds_describeDbinstancesCmd)
}
