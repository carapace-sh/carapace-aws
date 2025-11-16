package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Create a new cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_createClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_createClusterCmd).Standalone()

		route53RecoveryControlConfig_createClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters.")
		route53RecoveryControlConfig_createClusterCmd.Flags().String("cluster-name", "", "The name of the cluster.")
		route53RecoveryControlConfig_createClusterCmd.Flags().String("network-type", "", "The network type of the cluster.")
		route53RecoveryControlConfig_createClusterCmd.Flags().String("tags", "", "The tags associated with the cluster.")
		route53RecoveryControlConfig_createClusterCmd.MarkFlagRequired("cluster-name")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_createClusterCmd)
}
