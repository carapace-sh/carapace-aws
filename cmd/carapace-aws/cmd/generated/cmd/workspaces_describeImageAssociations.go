package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeImageAssociationsCmd = &cobra.Command{
	Use:   "describe-image-associations",
	Short: "Describes the associations between the applications and the specified image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeImageAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeImageAssociationsCmd).Standalone()

		workspaces_describeImageAssociationsCmd.Flags().String("associated-resource-types", "", "The resource types of the associated resource.")
		workspaces_describeImageAssociationsCmd.Flags().String("image-id", "", "The identifier of the image.")
		workspaces_describeImageAssociationsCmd.MarkFlagRequired("associated-resource-types")
		workspaces_describeImageAssociationsCmd.MarkFlagRequired("image-id")
	})
	workspacesCmd.AddCommand(workspaces_describeImageAssociationsCmd)
}
