package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxChangesetsCmd = &cobra.Command{
	Use:   "list-kx-changesets",
	Short: "Returns a list of all the changesets for a database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxChangesetsCmd).Standalone()

	finspace_listKxChangesetsCmd.Flags().String("database-name", "", "The name of the kdb database.")
	finspace_listKxChangesetsCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_listKxChangesetsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	finspace_listKxChangesetsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspace_listKxChangesetsCmd.MarkFlagRequired("database-name")
	finspace_listKxChangesetsCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_listKxChangesetsCmd)
}
