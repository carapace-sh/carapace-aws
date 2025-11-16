package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclusterEndpointsCmd = &cobra.Command{
	Use:   "describe-dbcluster-endpoints",
	Short: "Returns information about endpoints for an Amazon Aurora DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclusterEndpointsCmd).Standalone()

	rds_describeDbclusterEndpointsCmd.Flags().String("dbcluster-endpoint-identifier", "", "The identifier of the endpoint to describe.")
	rds_describeDbclusterEndpointsCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the DB cluster associated with the endpoint.")
	rds_describeDbclusterEndpointsCmd.Flags().String("filters", "", "A set of name-value pairs that define which endpoints to include in the output.")
	rds_describeDbclusterEndpointsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterEndpoints` request.")
	rds_describeDbclusterEndpointsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rdsCmd.AddCommand(rds_describeDbclusterEndpointsCmd)
}
