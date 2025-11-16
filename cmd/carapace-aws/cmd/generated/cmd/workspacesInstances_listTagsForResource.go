package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves tags for a WorkSpace Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesInstances_listTagsForResourceCmd).Standalone()

		workspacesInstances_listTagsForResourceCmd.Flags().String("workspace-instance-id", "", "Unique identifier of the WorkSpace Instance.")
		workspacesInstances_listTagsForResourceCmd.MarkFlagRequired("workspace-instance-id")
	})
	workspacesInstancesCmd.AddCommand(workspacesInstances_listTagsForResourceCmd)
}
