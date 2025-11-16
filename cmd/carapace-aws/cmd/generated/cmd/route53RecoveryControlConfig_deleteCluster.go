package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Delete a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_deleteClusterCmd).Standalone()

	route53RecoveryControlConfig_deleteClusterCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster that you're deleting.")
	route53RecoveryControlConfig_deleteClusterCmd.MarkFlagRequired("cluster-arn")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_deleteClusterCmd)
}
