package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_listSyncConfigurationsCmd = &cobra.Command{
	Use:   "list-sync-configurations",
	Short: "Returns a list of sync configurations for a specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_listSyncConfigurationsCmd).Standalone()

	codestarConnections_listSyncConfigurationsCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codestarConnections_listSyncConfigurationsCmd.Flags().String("next-token", "", "An enumeration token that allows the operation to batch the results of the operation.")
	codestarConnections_listSyncConfigurationsCmd.Flags().String("repository-link-id", "", "The ID of the repository link for the requested list of sync configurations.")
	codestarConnections_listSyncConfigurationsCmd.Flags().String("sync-type", "", "The sync type for the requested list of sync configurations.")
	codestarConnections_listSyncConfigurationsCmd.MarkFlagRequired("repository-link-id")
	codestarConnections_listSyncConfigurationsCmd.MarkFlagRequired("sync-type")
	codestarConnectionsCmd.AddCommand(codestarConnections_listSyncConfigurationsCmd)
}
