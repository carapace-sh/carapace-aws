package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_updateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Updates settings for a connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_updateConnectionCmd).Standalone()

	events_updateConnectionCmd.Flags().String("auth-parameters", "", "The authorization parameters to use for the connection.")
	events_updateConnectionCmd.Flags().String("authorization-type", "", "The type of authorization to use for the connection.")
	events_updateConnectionCmd.Flags().String("description", "", "A description for the connection.")
	events_updateConnectionCmd.Flags().String("invocation-connectivity-parameters", "", "For connections to private APIs, the parameters to use for invoking the API.")
	events_updateConnectionCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection.")
	events_updateConnectionCmd.Flags().String("name", "", "The name of the connection to update.")
	events_updateConnectionCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_updateConnectionCmd)
}
