package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createRelayCmd = &cobra.Command{
	Use:   "create-relay",
	Short: "Creates a relay resource which can be used in rules to relay incoming emails to defined relay destinations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createRelayCmd).Standalone()

	mailmanager_createRelayCmd.Flags().String("authentication", "", "Authentication for the relay destination server—specify the secretARN where the SMTP credentials are stored.")
	mailmanager_createRelayCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
	mailmanager_createRelayCmd.Flags().String("relay-name", "", "The unique name of the relay resource.")
	mailmanager_createRelayCmd.Flags().String("server-name", "", "The destination relay server address.")
	mailmanager_createRelayCmd.Flags().String("server-port", "", "The destination relay server port.")
	mailmanager_createRelayCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
	mailmanager_createRelayCmd.MarkFlagRequired("authentication")
	mailmanager_createRelayCmd.MarkFlagRequired("relay-name")
	mailmanager_createRelayCmd.MarkFlagRequired("server-name")
	mailmanager_createRelayCmd.MarkFlagRequired("server-port")
	mailmanagerCmd.AddCommand(mailmanager_createRelayCmd)
}
