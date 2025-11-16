package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteAppBlockBuilderCmd = &cobra.Command{
	Use:   "delete-app-block-builder",
	Short: "Deletes an app block builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteAppBlockBuilderCmd).Standalone()

	appstream_deleteAppBlockBuilderCmd.Flags().String("name", "", "The name of the app block builder.")
	appstream_deleteAppBlockBuilderCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_deleteAppBlockBuilderCmd)
}
