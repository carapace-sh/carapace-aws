package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_createResourceExplorerSetupCmd = &cobra.Command{
	Use:   "create-resource-explorer-setup",
	Short: "Creates a Resource Explorer setup configuration across multiple Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_createResourceExplorerSetupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_createResourceExplorerSetupCmd).Standalone()

		resourceExplorer2_createResourceExplorerSetupCmd.Flags().String("aggregator-regions", "", "A list of Amazon Web Services Regions that should be configured as aggregator Regions.")
		resourceExplorer2_createResourceExplorerSetupCmd.Flags().String("region-list", "", "A list of Amazon Web Services Regions where Resource Explorer should be configured.")
		resourceExplorer2_createResourceExplorerSetupCmd.Flags().String("view-name", "", "The name for the view to be created as part of the Resource Explorer setup.")
		resourceExplorer2_createResourceExplorerSetupCmd.MarkFlagRequired("region-list")
		resourceExplorer2_createResourceExplorerSetupCmd.MarkFlagRequired("view-name")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_createResourceExplorerSetupCmd)
}
