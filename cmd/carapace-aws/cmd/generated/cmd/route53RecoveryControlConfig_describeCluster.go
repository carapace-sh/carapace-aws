package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_describeClusterCmd = &cobra.Command{
	Use:   "describe-cluster",
	Short: "Display the details about a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_describeClusterCmd).Standalone()

	route53RecoveryControlConfig_describeClusterCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
	route53RecoveryControlConfig_describeClusterCmd.MarkFlagRequired("cluster-arn")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_describeClusterCmd)
}
