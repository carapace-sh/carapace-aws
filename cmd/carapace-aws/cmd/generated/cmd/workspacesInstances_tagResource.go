package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a WorkSpace Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_tagResourceCmd).Standalone()

	workspacesInstances_tagResourceCmd.Flags().String("tags", "", "Tags to be added to the WorkSpace Instance.")
	workspacesInstances_tagResourceCmd.Flags().String("workspace-instance-id", "", "Unique identifier of the WorkSpace Instance to tag.")
	workspacesInstances_tagResourceCmd.MarkFlagRequired("tags")
	workspacesInstances_tagResourceCmd.MarkFlagRequired("workspace-instance-id")
	workspacesInstancesCmd.AddCommand(workspacesInstances_tagResourceCmd)
}
