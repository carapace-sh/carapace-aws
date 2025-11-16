package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_batchDeleteImportDataCmd = &cobra.Command{
	Use:   "batch-delete-import-data",
	Short: "Deletes one or more import tasks, each identified by their import ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_batchDeleteImportDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_batchDeleteImportDataCmd).Standalone()

		discovery_batchDeleteImportDataCmd.Flags().Bool("delete-history", false, "Set to `true` to remove the deleted import task from [DescribeImportTasks]().")
		discovery_batchDeleteImportDataCmd.Flags().String("import-task-ids", "", "The IDs for the import tasks that you want to delete.")
		discovery_batchDeleteImportDataCmd.Flags().Bool("no-delete-history", false, "Set to `true` to remove the deleted import task from [DescribeImportTasks]().")
		discovery_batchDeleteImportDataCmd.MarkFlagRequired("import-task-ids")
		discovery_batchDeleteImportDataCmd.Flag("no-delete-history").Hidden = true
	})
	discoveryCmd.AddCommand(discovery_batchDeleteImportDataCmd)
}
