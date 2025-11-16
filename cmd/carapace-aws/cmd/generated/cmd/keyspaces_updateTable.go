package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_updateTableCmd = &cobra.Command{
	Use:   "update-table",
	Short: "Adds new columns to the table or updates one of the table's settings, for example capacity mode, auto scaling, encryption, point-in-time recovery, or ttl settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_updateTableCmd).Standalone()

	keyspaces_updateTableCmd.Flags().String("add-columns", "", "For each column to be added to the specified table:")
	keyspaces_updateTableCmd.Flags().String("auto-scaling-specification", "", "The optional auto scaling settings to update for a table in provisioned capacity mode.")
	keyspaces_updateTableCmd.Flags().String("capacity-specification", "", "Modifies the read/write throughput capacity mode for the table.")
	keyspaces_updateTableCmd.Flags().String("cdc-specification", "", "The CDC stream settings of the table.")
	keyspaces_updateTableCmd.Flags().String("client-side-timestamps", "", "Enables client-side timestamps for the table.")
	keyspaces_updateTableCmd.Flags().String("default-time-to-live", "", "The default Time to Live setting in seconds for the table.")
	keyspaces_updateTableCmd.Flags().String("encryption-specification", "", "Modifies the encryption settings of the table.")
	keyspaces_updateTableCmd.Flags().String("keyspace-name", "", "The name of the keyspace the specified table is stored in.")
	keyspaces_updateTableCmd.Flags().String("point-in-time-recovery", "", "Modifies the `pointInTimeRecovery` settings of the table.")
	keyspaces_updateTableCmd.Flags().String("replica-specifications", "", "The Region specific settings of a multi-Regional table.")
	keyspaces_updateTableCmd.Flags().String("table-name", "", "The name of the table.")
	keyspaces_updateTableCmd.Flags().String("ttl", "", "Modifies Time to Live custom settings for the table.")
	keyspaces_updateTableCmd.MarkFlagRequired("keyspace-name")
	keyspaces_updateTableCmd.MarkFlagRequired("table-name")
	keyspacesCmd.AddCommand(keyspaces_updateTableCmd)
}
