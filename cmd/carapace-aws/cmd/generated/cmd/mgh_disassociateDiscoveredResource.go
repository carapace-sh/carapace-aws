package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_disassociateDiscoveredResourceCmd = &cobra.Command{
	Use:   "disassociate-discovered-resource",
	Short: "Disassociate an Application Discovery Service discovered resource from a migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_disassociateDiscoveredResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_disassociateDiscoveredResourceCmd).Standalone()

		mgh_disassociateDiscoveredResourceCmd.Flags().String("configuration-id", "", "ConfigurationId of the Application Discovery Service resource to be disassociated.")
		mgh_disassociateDiscoveredResourceCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
		mgh_disassociateDiscoveredResourceCmd.Flags().String("migration-task-name", "", "The identifier given to the MigrationTask.")
		mgh_disassociateDiscoveredResourceCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
		mgh_disassociateDiscoveredResourceCmd.MarkFlagRequired("configuration-id")
		mgh_disassociateDiscoveredResourceCmd.MarkFlagRequired("migration-task-name")
		mgh_disassociateDiscoveredResourceCmd.MarkFlagRequired("progress-update-stream")
	})
	mghCmd.AddCommand(mgh_disassociateDiscoveredResourceCmd)
}
