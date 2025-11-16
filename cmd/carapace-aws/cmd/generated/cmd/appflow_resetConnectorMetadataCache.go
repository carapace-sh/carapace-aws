package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_resetConnectorMetadataCacheCmd = &cobra.Command{
	Use:   "reset-connector-metadata-cache",
	Short: "Resets metadata about your connector entities that Amazon AppFlow stored in its cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_resetConnectorMetadataCacheCmd).Standalone()

	appflow_resetConnectorMetadataCacheCmd.Flags().String("api-version", "", "The API version that you specified in the connector profile that you’re resetting cached metadata for.")
	appflow_resetConnectorMetadataCacheCmd.Flags().String("connector-entity-name", "", "Use this parameter if you want to reset cached metadata about the details for an individual entity.")
	appflow_resetConnectorMetadataCacheCmd.Flags().String("connector-profile-name", "", "The name of the connector profile that you want to reset cached metadata for.")
	appflow_resetConnectorMetadataCacheCmd.Flags().String("connector-type", "", "The type of connector to reset cached metadata for.")
	appflow_resetConnectorMetadataCacheCmd.Flags().String("entities-path", "", "Use this parameter only if you’re resetting the cached metadata about a nested entity.")
	appflowCmd.AddCommand(appflow_resetConnectorMetadataCacheCmd)
}
