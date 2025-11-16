package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbproxiesCmd = &cobra.Command{
	Use:   "describe-dbproxies",
	Short: "Returns information about DB proxies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbproxiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbproxiesCmd).Standalone()

		rds_describeDbproxiesCmd.Flags().String("dbproxy-name", "", "The name of the DB proxy.")
		rds_describeDbproxiesCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		rds_describeDbproxiesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		rds_describeDbproxiesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbproxiesCmd)
}
