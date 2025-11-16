package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "Retrieve a list of sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listSessionsCmd).Standalone()

		glue_listSessionsCmd.Flags().String("max-results", "", "The maximum number of results.")
		glue_listSessionsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more result.")
		glue_listSessionsCmd.Flags().String("request-origin", "", "The origin of the request.")
		glue_listSessionsCmd.Flags().String("tags", "", "Tags belonging to the session.")
	})
	glueCmd.AddCommand(glue_listSessionsCmd)
}
