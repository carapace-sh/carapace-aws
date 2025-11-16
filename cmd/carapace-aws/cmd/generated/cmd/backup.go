package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup\n\nBackup is a unified backup service designed to protect Amazon Web Services services and their associated data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupCmd).Standalone()

	rootCmd.AddCommand(backupCmd)
}
