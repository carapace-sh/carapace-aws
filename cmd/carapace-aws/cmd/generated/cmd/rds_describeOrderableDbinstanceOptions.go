package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeOrderableDbinstanceOptionsCmd = &cobra.Command{
	Use:   "describe-orderable-dbinstance-options",
	Short: "Describes the orderable DB instance options for a specified DB engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeOrderableDbinstanceOptionsCmd).Standalone()

	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("availability-zone-group", "", "The Availability Zone group associated with a Local Zone.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("dbinstance-class", "", "A filter to include only the available options for the specified DB instance class.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("engine", "", "The name of the database engine to describe DB instance options for.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("engine-version", "", "A filter to include only the available options for the specified engine version.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("license-model", "", "A filter to include only the available options for the specified license model.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions request.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeOrderableDbinstanceOptionsCmd.Flags().String("vpc", "", "Specifies whether to show only VPC or non-VPC offerings.")
	rds_describeOrderableDbinstanceOptionsCmd.MarkFlagRequired("engine")
	rdsCmd.AddCommand(rds_describeOrderableDbinstanceOptionsCmd)
}
