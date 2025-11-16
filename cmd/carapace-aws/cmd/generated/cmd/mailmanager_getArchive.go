package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getArchiveCmd = &cobra.Command{
	Use:   "get-archive",
	Short: "Retrieves the full details and current state of a specified email archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getArchiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getArchiveCmd).Standalone()

		mailmanager_getArchiveCmd.Flags().String("archive-id", "", "The identifier of the archive to retrieve.")
		mailmanager_getArchiveCmd.MarkFlagRequired("archive-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getArchiveCmd)
}
