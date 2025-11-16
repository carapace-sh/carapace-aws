package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteBillingGroupCmd = &cobra.Command{
	Use:   "delete-billing-group",
	Short: "Deletes the billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteBillingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteBillingGroupCmd).Standalone()

		iot_deleteBillingGroupCmd.Flags().String("billing-group-name", "", "The name of the billing group.")
		iot_deleteBillingGroupCmd.Flags().String("expected-version", "", "The expected version of the billing group.")
		iot_deleteBillingGroupCmd.MarkFlagRequired("billing-group-name")
	})
	iotCmd.AddCommand(iot_deleteBillingGroupCmd)
}
