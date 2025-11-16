package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getArchiveMessageCmd = &cobra.Command{
	Use:   "get-archive-message",
	Short: "Returns a pre-signed URL that provides temporary download access to the specific email message stored in the archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getArchiveMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getArchiveMessageCmd).Standalone()

		mailmanager_getArchiveMessageCmd.Flags().String("archived-message-id", "", "The unique identifier of the archived email message.")
		mailmanager_getArchiveMessageCmd.MarkFlagRequired("archived-message-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getArchiveMessageCmd)
}
