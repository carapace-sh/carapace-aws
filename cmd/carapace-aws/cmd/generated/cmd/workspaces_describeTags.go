package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Describes the specified tags for the specified WorkSpaces resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeTagsCmd).Standalone()

		workspaces_describeTagsCmd.Flags().String("resource-id", "", "The identifier of the WorkSpaces resource.")
		workspaces_describeTagsCmd.MarkFlagRequired("resource-id")
	})
	workspacesCmd.AddCommand(workspaces_describeTagsCmd)
}
