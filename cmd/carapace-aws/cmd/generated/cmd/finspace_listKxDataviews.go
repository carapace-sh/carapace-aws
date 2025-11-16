package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxDataviewsCmd = &cobra.Command{
	Use:   "list-kx-dataviews",
	Short: "Returns a list of all the dataviews in the database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxDataviewsCmd).Standalone()

	finspace_listKxDataviewsCmd.Flags().String("database-name", "", "The name of the database where the dataviews were created.")
	finspace_listKxDataviewsCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, for which you want to retrieve a list of dataviews.")
	finspace_listKxDataviewsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	finspace_listKxDataviewsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspace_listKxDataviewsCmd.MarkFlagRequired("database-name")
	finspace_listKxDataviewsCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_listKxDataviewsCmd)
}
