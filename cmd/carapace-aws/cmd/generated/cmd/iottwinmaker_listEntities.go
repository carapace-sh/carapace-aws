package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listEntitiesCmd = &cobra.Command{
	Use:   "list-entities",
	Short: "Lists all entities in a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listEntitiesCmd).Standalone()

	iottwinmaker_listEntitiesCmd.Flags().String("filters", "", "A list of objects that filter the request.")
	iottwinmaker_listEntitiesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iottwinmaker_listEntitiesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iottwinmaker_listEntitiesCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
	iottwinmaker_listEntitiesCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_listEntitiesCmd)
}
