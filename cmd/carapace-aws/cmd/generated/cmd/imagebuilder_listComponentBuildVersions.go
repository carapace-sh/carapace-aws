package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listComponentBuildVersionsCmd = &cobra.Command{
	Use:   "list-component-build-versions",
	Short: "Returns the list of component build versions for the specified component version Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listComponentBuildVersionsCmd).Standalone()

	imagebuilder_listComponentBuildVersionsCmd.Flags().String("component-version-arn", "", "The component version Amazon Resource Name (ARN) whose versions you want to list.")
	imagebuilder_listComponentBuildVersionsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listComponentBuildVersionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilderCmd.AddCommand(imagebuilder_listComponentBuildVersionsCmd)
}
