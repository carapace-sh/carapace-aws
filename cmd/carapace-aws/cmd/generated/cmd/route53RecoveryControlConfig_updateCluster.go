package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Updates an existing cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_updateClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_updateClusterCmd).Standalone()

		route53RecoveryControlConfig_updateClusterCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
		route53RecoveryControlConfig_updateClusterCmd.Flags().String("network-type", "", "The network type of the cluster.")
		route53RecoveryControlConfig_updateClusterCmd.MarkFlagRequired("cluster-arn")
		route53RecoveryControlConfig_updateClusterCmd.MarkFlagRequired("network-type")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_updateClusterCmd)
}
