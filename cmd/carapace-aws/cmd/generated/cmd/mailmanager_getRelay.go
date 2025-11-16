package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getRelayCmd = &cobra.Command{
	Use:   "get-relay",
	Short: "Fetch the relay resource and it's attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getRelayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getRelayCmd).Standalone()

		mailmanager_getRelayCmd.Flags().String("relay-id", "", "A unique relay identifier.")
		mailmanager_getRelayCmd.MarkFlagRequired("relay-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getRelayCmd)
}
