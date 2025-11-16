package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbproxyTargetGroupCmd = &cobra.Command{
	Use:   "modify-dbproxy-target-group",
	Short: "Modifies the properties of a `DBProxyTargetGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbproxyTargetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyDbproxyTargetGroupCmd).Standalone()

		rds_modifyDbproxyTargetGroupCmd.Flags().String("connection-pool-config", "", "The settings that determine the size and behavior of the connection pool for the target group.")
		rds_modifyDbproxyTargetGroupCmd.Flags().String("dbproxy-name", "", "The name of the proxy.")
		rds_modifyDbproxyTargetGroupCmd.Flags().String("new-name", "", "The new name for the modified `DBProxyTarget`.")
		rds_modifyDbproxyTargetGroupCmd.Flags().String("target-group-name", "", "The name of the target group to modify.")
		rds_modifyDbproxyTargetGroupCmd.MarkFlagRequired("dbproxy-name")
		rds_modifyDbproxyTargetGroupCmd.MarkFlagRequired("target-group-name")
	})
	rdsCmd.AddCommand(rds_modifyDbproxyTargetGroupCmd)
}
