package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getComponentCmd = &cobra.Command{
	Use:   "get-component",
	Short: "Gets a component object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_getComponentCmd).Standalone()

		imagebuilder_getComponentCmd.Flags().String("component-build-version-arn", "", "The Amazon Resource Name (ARN) of the component that you want to get.")
		imagebuilder_getComponentCmd.MarkFlagRequired("component-build-version-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_getComponentCmd)
}
