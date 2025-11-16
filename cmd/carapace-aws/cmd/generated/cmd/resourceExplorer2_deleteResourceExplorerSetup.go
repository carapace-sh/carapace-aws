package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_deleteResourceExplorerSetupCmd = &cobra.Command{
	Use:   "delete-resource-explorer-setup",
	Short: "Deletes a Resource Explorer setup configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_deleteResourceExplorerSetupCmd).Standalone()

	resourceExplorer2_deleteResourceExplorerSetupCmd.Flags().Bool("delete-in-all-regions", false, "Specifies whether to delete Resource Explorer configuration from all Regions where it is currently enabled.")
	resourceExplorer2_deleteResourceExplorerSetupCmd.Flags().Bool("no-delete-in-all-regions", false, "Specifies whether to delete Resource Explorer configuration from all Regions where it is currently enabled.")
	resourceExplorer2_deleteResourceExplorerSetupCmd.Flags().String("region-list", "", "A list of Amazon Web Services Regions from which to delete the Resource Explorer configuration.")
	resourceExplorer2_deleteResourceExplorerSetupCmd.Flag("no-delete-in-all-regions").Hidden = true
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_deleteResourceExplorerSetupCmd)
}
