package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_uploadEntityDefinitionsCmd = &cobra.Command{
	Use:   "upload-entity-definitions",
	Short: "Asynchronously uploads one or more entity definitions to the user's namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_uploadEntityDefinitionsCmd).Standalone()

	iotthingsgraph_uploadEntityDefinitionsCmd.Flags().String("deprecate-existing-entities", "", "A Boolean that specifies whether to deprecate all entities in the latest version before uploading the new `DefinitionDocument`.")
	iotthingsgraph_uploadEntityDefinitionsCmd.Flags().String("document", "", "The `DefinitionDocument` that defines the updated entities.")
	iotthingsgraph_uploadEntityDefinitionsCmd.Flags().String("sync-with-public-namespace", "", "A Boolean that specifies whether to synchronize with the latest version of the public namespace.")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_uploadEntityDefinitionsCmd)
}
