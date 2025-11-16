package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_deleteRoutingControlCmd = &cobra.Command{
	Use:   "delete-routing-control",
	Short: "Deletes a routing control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_deleteRoutingControlCmd).Standalone()

	route53RecoveryControlConfig_deleteRoutingControlCmd.Flags().String("routing-control-arn", "", "The Amazon Resource Name (ARN) of the routing control that you're deleting.")
	route53RecoveryControlConfig_deleteRoutingControlCmd.MarkFlagRequired("routing-control-arn")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_deleteRoutingControlCmd)
}
