package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getEntityCmd = &cobra.Command{
	Use:   "get-entity",
	Short: "Retrieves information about an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getEntityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_getEntityCmd).Standalone()

		iottwinmaker_getEntityCmd.Flags().String("entity-id", "", "The ID of the entity.")
		iottwinmaker_getEntityCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
		iottwinmaker_getEntityCmd.MarkFlagRequired("entity-id")
		iottwinmaker_getEntityCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_getEntityCmd)
}
