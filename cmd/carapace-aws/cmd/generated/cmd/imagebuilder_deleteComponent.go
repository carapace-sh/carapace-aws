package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteComponentCmd = &cobra.Command{
	Use:   "delete-component",
	Short: "Deletes a component build version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_deleteComponentCmd).Standalone()

		imagebuilder_deleteComponentCmd.Flags().String("component-build-version-arn", "", "The Amazon Resource Name (ARN) of the component build version to delete.")
		imagebuilder_deleteComponentCmd.MarkFlagRequired("component-build-version-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_deleteComponentCmd)
}
