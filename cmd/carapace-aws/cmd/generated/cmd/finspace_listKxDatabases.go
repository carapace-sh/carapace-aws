package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxDatabasesCmd = &cobra.Command{
	Use:   "list-kx-databases",
	Short: "Returns a list of all the databases in the kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxDatabasesCmd).Standalone()

	finspace_listKxDatabasesCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_listKxDatabasesCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	finspace_listKxDatabasesCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspace_listKxDatabasesCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_listKxDatabasesCmd)
}
