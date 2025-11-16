package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeCopyJobCmd = &cobra.Command{
	Use:   "describe-copy-job",
	Short: "Returns metadata associated with creating a copy of a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeCopyJobCmd).Standalone()

	backup_describeCopyJobCmd.Flags().String("copy-job-id", "", "Uniquely identifies a copy job.")
	backup_describeCopyJobCmd.MarkFlagRequired("copy-job-id")
	backupCmd.AddCommand(backup_describeCopyJobCmd)
}
