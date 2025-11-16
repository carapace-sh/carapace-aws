package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateBillingGroupCmd = &cobra.Command{
	Use:   "update-billing-group",
	Short: "Updates information about the billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateBillingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateBillingGroupCmd).Standalone()

		iot_updateBillingGroupCmd.Flags().String("billing-group-name", "", "The name of the billing group.")
		iot_updateBillingGroupCmd.Flags().String("billing-group-properties", "", "The properties of the billing group.")
		iot_updateBillingGroupCmd.Flags().String("expected-version", "", "The expected version of the billing group.")
		iot_updateBillingGroupCmd.MarkFlagRequired("billing-group-name")
		iot_updateBillingGroupCmd.MarkFlagRequired("billing-group-properties")
	})
	iotCmd.AddCommand(iot_updateBillingGroupCmd)
}
