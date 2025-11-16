package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listSchemasCmd = &cobra.Command{
	Use:   "list-schemas",
	Short: "Returns the list of schemas associated with the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listSchemasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listSchemasCmd).Standalone()

		personalize_listSchemasCmd.Flags().String("max-results", "", "The maximum number of schemas to return.")
		personalize_listSchemasCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListSchemas` for getting the next set of schemas (if they exist).")
	})
	personalizeCmd.AddCommand(personalize_listSchemasCmd)
}
