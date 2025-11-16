package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbproxyEndpointsCmd = &cobra.Command{
	Use:   "describe-dbproxy-endpoints",
	Short: "Returns information about DB proxy endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbproxyEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbproxyEndpointsCmd).Standalone()

		rds_describeDbproxyEndpointsCmd.Flags().String("dbproxy-endpoint-name", "", "The name of a DB proxy endpoint to describe.")
		rds_describeDbproxyEndpointsCmd.Flags().String("dbproxy-name", "", "The name of the DB proxy whose endpoints you want to describe.")
		rds_describeDbproxyEndpointsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		rds_describeDbproxyEndpointsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		rds_describeDbproxyEndpointsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbproxyEndpointsCmd)
}
