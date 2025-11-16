package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getArchiveMessageContentCmd = &cobra.Command{
	Use:   "get-archive-message-content",
	Short: "Returns the textual content of a specific email message stored in the archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getArchiveMessageContentCmd).Standalone()

	mailmanager_getArchiveMessageContentCmd.Flags().String("archived-message-id", "", "The unique identifier of the archived email message.")
	mailmanager_getArchiveMessageContentCmd.MarkFlagRequired("archived-message-id")
	mailmanagerCmd.AddCommand(mailmanager_getArchiveMessageContentCmd)
}
