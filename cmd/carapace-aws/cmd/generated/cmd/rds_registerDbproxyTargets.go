package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_registerDbproxyTargetsCmd = &cobra.Command{
	Use:   "register-dbproxy-targets",
	Short: "Associate one or more `DBProxyTarget` data structures with a `DBProxyTargetGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_registerDbproxyTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_registerDbproxyTargetsCmd).Standalone()

		rds_registerDbproxyTargetsCmd.Flags().String("dbcluster-identifiers", "", "One or more DB cluster identifiers.")
		rds_registerDbproxyTargetsCmd.Flags().String("dbinstance-identifiers", "", "One or more DB instance identifiers.")
		rds_registerDbproxyTargetsCmd.Flags().String("dbproxy-name", "", "The identifier of the `DBProxy` that is associated with the `DBProxyTargetGroup`.")
		rds_registerDbproxyTargetsCmd.Flags().String("target-group-name", "", "The identifier of the `DBProxyTargetGroup`.")
		rds_registerDbproxyTargetsCmd.MarkFlagRequired("dbproxy-name")
	})
	rdsCmd.AddCommand(rds_registerDbproxyTargetsCmd)
}
