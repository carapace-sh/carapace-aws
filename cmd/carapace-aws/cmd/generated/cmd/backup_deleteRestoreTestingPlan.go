package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteRestoreTestingPlanCmd = &cobra.Command{
	Use:   "delete-restore-testing-plan",
	Short: "This request deletes the specified restore testing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteRestoreTestingPlanCmd).Standalone()

	backup_deleteRestoreTestingPlanCmd.Flags().String("restore-testing-plan-name", "", "Required unique name of the restore testing plan you wish to delete.")
	backup_deleteRestoreTestingPlanCmd.MarkFlagRequired("restore-testing-plan-name")
	backupCmd.AddCommand(backup_deleteRestoreTestingPlanCmd)
}
