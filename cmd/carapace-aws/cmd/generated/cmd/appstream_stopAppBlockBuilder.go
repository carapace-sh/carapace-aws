package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_stopAppBlockBuilderCmd = &cobra.Command{
	Use:   "stop-app-block-builder",
	Short: "Stops an app block builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_stopAppBlockBuilderCmd).Standalone()

	appstream_stopAppBlockBuilderCmd.Flags().String("name", "", "The name of the app block builder.")
	appstream_stopAppBlockBuilderCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_stopAppBlockBuilderCmd)
}
