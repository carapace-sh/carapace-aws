package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_reloadReplicationTablesCmd = &cobra.Command{
	Use:   "reload-replication-tables",
	Short: "Reloads the target database table with the source data for a given DMS Serverless replication configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_reloadReplicationTablesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_reloadReplicationTablesCmd).Standalone()

		dms_reloadReplicationTablesCmd.Flags().String("reload-option", "", "Options for reload.")
		dms_reloadReplicationTablesCmd.Flags().String("replication-config-arn", "", "The Amazon Resource Name of the replication config for which to reload tables.")
		dms_reloadReplicationTablesCmd.Flags().String("tables-to-reload", "", "The list of tables to reload.")
		dms_reloadReplicationTablesCmd.MarkFlagRequired("replication-config-arn")
		dms_reloadReplicationTablesCmd.MarkFlagRequired("tables-to-reload")
	})
	dmsCmd.AddCommand(dms_reloadReplicationTablesCmd)
}
