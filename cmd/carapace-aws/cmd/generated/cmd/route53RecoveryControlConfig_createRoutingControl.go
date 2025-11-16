package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_createRoutingControlCmd = &cobra.Command{
	Use:   "create-routing-control",
	Short: "Creates a new routing control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_createRoutingControlCmd).Standalone()

	route53RecoveryControlConfig_createRoutingControlCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters.")
	route53RecoveryControlConfig_createRoutingControlCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster that includes the routing control.")
	route53RecoveryControlConfig_createRoutingControlCmd.Flags().String("control-panel-arn", "", "The Amazon Resource Name (ARN) of the control panel that includes the routing control.")
	route53RecoveryControlConfig_createRoutingControlCmd.Flags().String("routing-control-name", "", "The name of the routing control.")
	route53RecoveryControlConfig_createRoutingControlCmd.MarkFlagRequired("cluster-arn")
	route53RecoveryControlConfig_createRoutingControlCmd.MarkFlagRequired("routing-control-name")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_createRoutingControlCmd)
}
