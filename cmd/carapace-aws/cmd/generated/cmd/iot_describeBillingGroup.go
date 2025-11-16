package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeBillingGroupCmd = &cobra.Command{
	Use:   "describe-billing-group",
	Short: "Returns information about a billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeBillingGroupCmd).Standalone()

	iot_describeBillingGroupCmd.Flags().String("billing-group-name", "", "The name of the billing group.")
	iot_describeBillingGroupCmd.MarkFlagRequired("billing-group-name")
	iotCmd.AddCommand(iot_describeBillingGroupCmd)
}
