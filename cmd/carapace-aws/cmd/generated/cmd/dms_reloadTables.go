package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_reloadTablesCmd = &cobra.Command{
	Use:   "reload-tables",
	Short: "Reloads the target database table with the source data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_reloadTablesCmd).Standalone()

	dms_reloadTablesCmd.Flags().String("reload-option", "", "Options for reload.")
	dms_reloadTablesCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) of the replication task.")
	dms_reloadTablesCmd.Flags().String("tables-to-reload", "", "The name and schema of the table to be reloaded.")
	dms_reloadTablesCmd.MarkFlagRequired("replication-task-arn")
	dms_reloadTablesCmd.MarkFlagRequired("tables-to-reload")
	dmsCmd.AddCommand(dms_reloadTablesCmd)
}
