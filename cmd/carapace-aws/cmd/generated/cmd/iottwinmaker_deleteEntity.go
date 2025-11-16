package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_deleteEntityCmd = &cobra.Command{
	Use:   "delete-entity",
	Short: "Deletes an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_deleteEntityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_deleteEntityCmd).Standalone()

		iottwinmaker_deleteEntityCmd.Flags().String("entity-id", "", "The ID of the entity to delete.")
		iottwinmaker_deleteEntityCmd.Flags().Bool("is-recursive", false, "A Boolean value that specifies whether the operation deletes child entities.")
		iottwinmaker_deleteEntityCmd.Flags().Bool("no-is-recursive", false, "A Boolean value that specifies whether the operation deletes child entities.")
		iottwinmaker_deleteEntityCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the entity to delete.")
		iottwinmaker_deleteEntityCmd.MarkFlagRequired("entity-id")
		iottwinmaker_deleteEntityCmd.Flag("no-is-recursive").Hidden = true
		iottwinmaker_deleteEntityCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_deleteEntityCmd)
}
