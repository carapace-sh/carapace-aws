package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteRelayCmd = &cobra.Command{
	Use:   "delete-relay",
	Short: "Deletes an existing relay resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteRelayCmd).Standalone()

	mailmanager_deleteRelayCmd.Flags().String("relay-id", "", "The unique relay identifier.")
	mailmanager_deleteRelayCmd.MarkFlagRequired("relay-id")
	mailmanagerCmd.AddCommand(mailmanager_deleteRelayCmd)
}
