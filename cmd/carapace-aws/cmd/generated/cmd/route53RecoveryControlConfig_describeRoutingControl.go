package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_describeRoutingControlCmd = &cobra.Command{
	Use:   "describe-routing-control",
	Short: "Displays details about a routing control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_describeRoutingControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_describeRoutingControlCmd).Standalone()

		route53RecoveryControlConfig_describeRoutingControlCmd.Flags().String("routing-control-arn", "", "The Amazon Resource Name (ARN) of the routing control.")
		route53RecoveryControlConfig_describeRoutingControlCmd.MarkFlagRequired("routing-control-arn")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_describeRoutingControlCmd)
}
