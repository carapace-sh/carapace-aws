package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_startAppBlockBuilderCmd = &cobra.Command{
	Use:   "start-app-block-builder",
	Short: "Starts an app block builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_startAppBlockBuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_startAppBlockBuilderCmd).Standalone()

		appstream_startAppBlockBuilderCmd.Flags().String("name", "", "The name of the app block builder.")
		appstream_startAppBlockBuilderCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_startAppBlockBuilderCmd)
}
