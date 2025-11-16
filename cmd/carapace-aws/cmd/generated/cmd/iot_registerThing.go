package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_registerThingCmd = &cobra.Command{
	Use:   "register-thing",
	Short: "Provisions a thing in the device registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_registerThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_registerThingCmd).Standalone()

		iot_registerThingCmd.Flags().String("parameters", "", "The parameters for provisioning a thing.")
		iot_registerThingCmd.Flags().String("template-body", "", "The provisioning template.")
		iot_registerThingCmd.MarkFlagRequired("template-body")
	})
	iotCmd.AddCommand(iot_registerThingCmd)
}
