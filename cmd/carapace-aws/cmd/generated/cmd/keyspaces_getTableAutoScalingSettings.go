package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_getTableAutoScalingSettingsCmd = &cobra.Command{
	Use:   "get-table-auto-scaling-settings",
	Short: "Returns auto scaling related settings of the specified table in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_getTableAutoScalingSettingsCmd).Standalone()

	keyspaces_getTableAutoScalingSettingsCmd.Flags().String("keyspace-name", "", "The name of the keyspace.")
	keyspaces_getTableAutoScalingSettingsCmd.Flags().String("table-name", "", "The name of the table.")
	keyspaces_getTableAutoScalingSettingsCmd.MarkFlagRequired("keyspace-name")
	keyspaces_getTableAutoScalingSettingsCmd.MarkFlagRequired("table-name")
	keyspacesCmd.AddCommand(keyspaces_getTableAutoScalingSettingsCmd)
}
