package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspaceBundlesCmd = &cobra.Command{
	Use:   "describe-workspace-bundles",
	Short: "Retrieves a list that describes the available WorkSpace bundles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspaceBundlesCmd).Standalone()

	workspaces_describeWorkspaceBundlesCmd.Flags().String("bundle-ids", "", "The identifiers of the bundles.")
	workspaces_describeWorkspaceBundlesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	workspaces_describeWorkspaceBundlesCmd.Flags().String("owner", "", "The owner of the bundles.")
	workspacesCmd.AddCommand(workspaces_describeWorkspaceBundlesCmd)
}
