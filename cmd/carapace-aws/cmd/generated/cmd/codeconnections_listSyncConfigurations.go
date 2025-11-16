package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_listSyncConfigurationsCmd = &cobra.Command{
	Use:   "list-sync-configurations",
	Short: "Returns a list of sync configurations for a specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_listSyncConfigurationsCmd).Standalone()

	codeconnections_listSyncConfigurationsCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codeconnections_listSyncConfigurationsCmd.Flags().String("next-token", "", "An enumeration token that allows the operation to batch the results of the operation.")
	codeconnections_listSyncConfigurationsCmd.Flags().String("repository-link-id", "", "The ID of the repository link for the requested list of sync configurations.")
	codeconnections_listSyncConfigurationsCmd.Flags().String("sync-type", "", "The sync type for the requested list of sync configurations.")
	codeconnections_listSyncConfigurationsCmd.MarkFlagRequired("repository-link-id")
	codeconnections_listSyncConfigurationsCmd.MarkFlagRequired("sync-type")
	codeconnectionsCmd.AddCommand(codeconnections_listSyncConfigurationsCmd)
}
