package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Returns an array of all the clusters in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_listClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_listClustersCmd).Standalone()

		route53RecoveryControlConfig_listClustersCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryControlConfig_listClustersCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_listClustersCmd)
}
