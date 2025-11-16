package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeBundleAssociationsCmd = &cobra.Command{
	Use:   "describe-bundle-associations",
	Short: "Describes the associations between the applications and the specified bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeBundleAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeBundleAssociationsCmd).Standalone()

		workspaces_describeBundleAssociationsCmd.Flags().String("associated-resource-types", "", "The resource types of the associated resource.")
		workspaces_describeBundleAssociationsCmd.Flags().String("bundle-id", "", "The identifier of the bundle.")
		workspaces_describeBundleAssociationsCmd.MarkFlagRequired("associated-resource-types")
		workspaces_describeBundleAssociationsCmd.MarkFlagRequired("bundle-id")
	})
	workspacesCmd.AddCommand(workspaces_describeBundleAssociationsCmd)
}
