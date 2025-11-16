package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_getMetadataCmd = &cobra.Command{
	Use:   "get-metadata",
	Short: "Returns existing metadata for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_getMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_getMetadataCmd).Standalone()

		amplifyuibuilder_getMetadataCmd.Flags().String("app-id", "", "The unique ID of the Amplify app.")
		amplifyuibuilder_getMetadataCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
		amplifyuibuilder_getMetadataCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_getMetadataCmd.MarkFlagRequired("environment-name")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_getMetadataCmd)
}
