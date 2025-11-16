package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_putMetadataFlagCmd = &cobra.Command{
	Use:   "put-metadata-flag",
	Short: "Stores the metadata information about a feature on a form.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_putMetadataFlagCmd).Standalone()

	amplifyuibuilder_putMetadataFlagCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
	amplifyuibuilder_putMetadataFlagCmd.Flags().String("body", "", "The metadata information to store.")
	amplifyuibuilder_putMetadataFlagCmd.Flags().String("environment-name", "", "The name of the backend environment that is part of the Amplify app.")
	amplifyuibuilder_putMetadataFlagCmd.Flags().String("feature-name", "", "The name of the feature associated with the metadata.")
	amplifyuibuilder_putMetadataFlagCmd.MarkFlagRequired("app-id")
	amplifyuibuilder_putMetadataFlagCmd.MarkFlagRequired("body")
	amplifyuibuilder_putMetadataFlagCmd.MarkFlagRequired("environment-name")
	amplifyuibuilder_putMetadataFlagCmd.MarkFlagRequired("feature-name")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_putMetadataFlagCmd)
}
