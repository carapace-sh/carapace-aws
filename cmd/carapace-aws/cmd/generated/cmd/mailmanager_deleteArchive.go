package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteArchiveCmd = &cobra.Command{
	Use:   "delete-archive",
	Short: "Initiates deletion of an email archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteArchiveCmd).Standalone()

	mailmanager_deleteArchiveCmd.Flags().String("archive-id", "", "The identifier of the archive to delete.")
	mailmanager_deleteArchiveCmd.MarkFlagRequired("archive-id")
	mailmanagerCmd.AddCommand(mailmanager_deleteArchiveCmd)
}
