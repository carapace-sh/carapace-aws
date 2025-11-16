package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_attachThingPrincipalCmd = &cobra.Command{
	Use:   "attach-thing-principal",
	Short: "Attaches the specified principal to the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_attachThingPrincipalCmd).Standalone()

	iot_attachThingPrincipalCmd.Flags().String("principal", "", "The principal, which can be a certificate ARN (as returned from the CreateCertificate operation) or an Amazon Cognito ID.")
	iot_attachThingPrincipalCmd.Flags().String("thing-name", "", "The name of the thing.")
	iot_attachThingPrincipalCmd.Flags().String("thing-principal-type", "", "The type of the relation you want to specify when you attach a principal to a thing.")
	iot_attachThingPrincipalCmd.MarkFlagRequired("principal")
	iot_attachThingPrincipalCmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_attachThingPrincipalCmd)
}
