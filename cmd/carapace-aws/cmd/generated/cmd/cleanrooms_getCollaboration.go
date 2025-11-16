package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getCollaborationCmd = &cobra.Command{
	Use:   "get-collaboration",
	Short: "Returns metadata about a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getCollaborationCmd).Standalone()

	cleanrooms_getCollaborationCmd.Flags().String("collaboration-identifier", "", "The identifier for the collaboration.")
	cleanrooms_getCollaborationCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getCollaborationCmd)
}
