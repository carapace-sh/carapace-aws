package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listCollaborationIdNamespaceAssociationsCmd = &cobra.Command{
	Use:   "list-collaboration-id-namespace-associations",
	Short: "Returns a list of the ID namespace associations in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listCollaborationIdNamespaceAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listCollaborationIdNamespaceAssociationsCmd).Standalone()

		cleanrooms_listCollaborationIdNamespaceAssociationsCmd.Flags().String("collaboration-identifier", "", "The unique identifier of the collaboration that contains the ID namespace associations that you want to retrieve.")
		cleanrooms_listCollaborationIdNamespaceAssociationsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
		cleanrooms_listCollaborationIdNamespaceAssociationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listCollaborationIdNamespaceAssociationsCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listCollaborationIdNamespaceAssociationsCmd)
}
