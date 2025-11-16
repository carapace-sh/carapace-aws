package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbproxyTargetsCmd = &cobra.Command{
	Use:   "describe-dbproxy-targets",
	Short: "Returns information about `DBProxyTarget` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbproxyTargetsCmd).Standalone()

	rds_describeDbproxyTargetsCmd.Flags().String("dbproxy-name", "", "The identifier of the `DBProxyTarget` to describe.")
	rds_describeDbproxyTargetsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	rds_describeDbproxyTargetsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	rds_describeDbproxyTargetsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeDbproxyTargetsCmd.Flags().String("target-group-name", "", "The identifier of the `DBProxyTargetGroup` to describe.")
	rds_describeDbproxyTargetsCmd.MarkFlagRequired("dbproxy-name")
	rdsCmd.AddCommand(rds_describeDbproxyTargetsCmd)
}
