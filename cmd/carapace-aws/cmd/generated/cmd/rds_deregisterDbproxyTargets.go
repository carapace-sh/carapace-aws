package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deregisterDbproxyTargetsCmd = &cobra.Command{
	Use:   "deregister-dbproxy-targets",
	Short: "Remove the association between one or more `DBProxyTarget` data structures and a `DBProxyTargetGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deregisterDbproxyTargetsCmd).Standalone()

	rds_deregisterDbproxyTargetsCmd.Flags().String("dbcluster-identifiers", "", "One or more DB cluster identifiers.")
	rds_deregisterDbproxyTargetsCmd.Flags().String("dbinstance-identifiers", "", "One or more DB instance identifiers.")
	rds_deregisterDbproxyTargetsCmd.Flags().String("dbproxy-name", "", "The identifier of the `DBProxy` that is associated with the `DBProxyTargetGroup`.")
	rds_deregisterDbproxyTargetsCmd.Flags().String("target-group-name", "", "The identifier of the `DBProxyTargetGroup`.")
	rds_deregisterDbproxyTargetsCmd.MarkFlagRequired("dbproxy-name")
	rdsCmd.AddCommand(rds_deregisterDbproxyTargetsCmd)
}
