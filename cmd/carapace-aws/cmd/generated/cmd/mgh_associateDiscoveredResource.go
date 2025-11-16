package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_associateDiscoveredResourceCmd = &cobra.Command{
	Use:   "associate-discovered-resource",
	Short: "Associates a discovered resource ID from Application Discovery Service with a migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_associateDiscoveredResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_associateDiscoveredResourceCmd).Standalone()

		mgh_associateDiscoveredResourceCmd.Flags().String("discovered-resource", "", "Object representing a Resource.")
		mgh_associateDiscoveredResourceCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
		mgh_associateDiscoveredResourceCmd.Flags().String("migration-task-name", "", "The identifier given to the MigrationTask.")
		mgh_associateDiscoveredResourceCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
		mgh_associateDiscoveredResourceCmd.MarkFlagRequired("discovered-resource")
		mgh_associateDiscoveredResourceCmd.MarkFlagRequired("migration-task-name")
		mgh_associateDiscoveredResourceCmd.MarkFlagRequired("progress-update-stream")
	})
	mghCmd.AddCommand(mgh_associateDiscoveredResourceCmd)
}
