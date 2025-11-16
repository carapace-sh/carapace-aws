package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_detachThingPrincipalCmd = &cobra.Command{
	Use:   "detach-thing-principal",
	Short: "Detaches the specified principal from the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_detachThingPrincipalCmd).Standalone()

	iot_detachThingPrincipalCmd.Flags().String("principal", "", "If the principal is a certificate, this value must be ARN of the certificate.")
	iot_detachThingPrincipalCmd.Flags().String("thing-name", "", "The name of the thing.")
	iot_detachThingPrincipalCmd.MarkFlagRequired("principal")
	iot_detachThingPrincipalCmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_detachThingPrincipalCmd)
}
