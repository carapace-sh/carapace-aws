package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_updateRelayCmd = &cobra.Command{
	Use:   "update-relay",
	Short: "Updates the attributes of an existing relay resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_updateRelayCmd).Standalone()

	mailmanager_updateRelayCmd.Flags().String("authentication", "", "Authentication for the relay destination server—specify the secretARN where the SMTP credentials are stored.")
	mailmanager_updateRelayCmd.Flags().String("relay-id", "", "The unique relay identifier.")
	mailmanager_updateRelayCmd.Flags().String("relay-name", "", "The name of the relay resource.")
	mailmanager_updateRelayCmd.Flags().String("server-name", "", "The destination relay server address.")
	mailmanager_updateRelayCmd.Flags().String("server-port", "", "The destination relay server port.")
	mailmanager_updateRelayCmd.MarkFlagRequired("relay-id")
	mailmanagerCmd.AddCommand(mailmanager_updateRelayCmd)
}
