package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_updateArchiveCmd = &cobra.Command{
	Use:   "update-archive",
	Short: "Updates the attributes of an existing email archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_updateArchiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_updateArchiveCmd).Standalone()

		mailmanager_updateArchiveCmd.Flags().String("archive-id", "", "The identifier of the archive to update.")
		mailmanager_updateArchiveCmd.Flags().String("archive-name", "", "A new, unique name for the archive.")
		mailmanager_updateArchiveCmd.Flags().String("retention", "", "A new retention period for emails in the archive.")
		mailmanager_updateArchiveCmd.MarkFlagRequired("archive-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_updateArchiveCmd)
}
