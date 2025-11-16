package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbproxyTargetGroupsCmd = &cobra.Command{
	Use:   "describe-dbproxy-target-groups",
	Short: "Returns information about DB proxy target groups, represented by `DBProxyTargetGroup` data structures.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbproxyTargetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbproxyTargetGroupsCmd).Standalone()

		rds_describeDbproxyTargetGroupsCmd.Flags().String("dbproxy-name", "", "The identifier of the `DBProxy` associated with the target group.")
		rds_describeDbproxyTargetGroupsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		rds_describeDbproxyTargetGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		rds_describeDbproxyTargetGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeDbproxyTargetGroupsCmd.Flags().String("target-group-name", "", "The identifier of the `DBProxyTargetGroup` to describe.")
		rds_describeDbproxyTargetGroupsCmd.MarkFlagRequired("dbproxy-name")
	})
	rdsCmd.AddCommand(rds_describeDbproxyTargetGroupsCmd)
}
