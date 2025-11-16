package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a WorkSpace Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_untagResourceCmd).Standalone()

	workspacesInstances_untagResourceCmd.Flags().String("tag-keys", "", "Keys of tags to be removed.")
	workspacesInstances_untagResourceCmd.Flags().String("workspace-instance-id", "", "Unique identifier of the WorkSpace Instance to untag.")
	workspacesInstances_untagResourceCmd.MarkFlagRequired("tag-keys")
	workspacesInstances_untagResourceCmd.MarkFlagRequired("workspace-instance-id")
	workspacesInstancesCmd.AddCommand(workspacesInstances_untagResourceCmd)
}
