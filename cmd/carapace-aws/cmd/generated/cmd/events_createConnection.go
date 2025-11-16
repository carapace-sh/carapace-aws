package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Creates a connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_createConnectionCmd).Standalone()

	events_createConnectionCmd.Flags().String("auth-parameters", "", "The authorization parameters to use to authorize with the endpoint.")
	events_createConnectionCmd.Flags().String("authorization-type", "", "The type of authorization to use for the connection.")
	events_createConnectionCmd.Flags().String("description", "", "A description for the connection to create.")
	events_createConnectionCmd.Flags().String("invocation-connectivity-parameters", "", "For connections to private APIs, the parameters to use for invoking the API.")
	events_createConnectionCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection.")
	events_createConnectionCmd.Flags().String("name", "", "The name for the connection to create.")
	events_createConnectionCmd.MarkFlagRequired("auth-parameters")
	events_createConnectionCmd.MarkFlagRequired("authorization-type")
	events_createConnectionCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_createConnectionCmd)
}
