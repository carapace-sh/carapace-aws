package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listIdNamespaceAssociationsCmd = &cobra.Command{
	Use:   "list-id-namespace-associations",
	Short: "Returns a list of ID namespace associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listIdNamespaceAssociationsCmd).Standalone()

	cleanrooms_listIdNamespaceAssociationsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanrooms_listIdNamespaceAssociationsCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID namespace association that you want to view.")
	cleanrooms_listIdNamespaceAssociationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listIdNamespaceAssociationsCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_listIdNamespaceAssociationsCmd)
}
