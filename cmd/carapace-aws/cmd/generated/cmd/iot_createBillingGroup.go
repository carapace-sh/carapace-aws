package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createBillingGroupCmd = &cobra.Command{
	Use:   "create-billing-group",
	Short: "Creates a billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createBillingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createBillingGroupCmd).Standalone()

		iot_createBillingGroupCmd.Flags().String("billing-group-name", "", "The name you wish to give to the billing group.")
		iot_createBillingGroupCmd.Flags().String("billing-group-properties", "", "The properties of the billing group.")
		iot_createBillingGroupCmd.Flags().String("tags", "", "Metadata which can be used to manage the billing group.")
		iot_createBillingGroupCmd.MarkFlagRequired("billing-group-name")
	})
	iotCmd.AddCommand(iot_createBillingGroupCmd)
}
