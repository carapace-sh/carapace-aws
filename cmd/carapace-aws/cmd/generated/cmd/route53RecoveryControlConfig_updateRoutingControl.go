package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_updateRoutingControlCmd = &cobra.Command{
	Use:   "update-routing-control",
	Short: "Updates a routing control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_updateRoutingControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_updateRoutingControlCmd).Standalone()

		route53RecoveryControlConfig_updateRoutingControlCmd.Flags().String("routing-control-arn", "", "The Amazon Resource Name (ARN) of the routing control.")
		route53RecoveryControlConfig_updateRoutingControlCmd.Flags().String("routing-control-name", "", "The name of the routing control.")
		route53RecoveryControlConfig_updateRoutingControlCmd.MarkFlagRequired("routing-control-arn")
		route53RecoveryControlConfig_updateRoutingControlCmd.MarkFlagRequired("routing-control-name")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_updateRoutingControlCmd)
}
