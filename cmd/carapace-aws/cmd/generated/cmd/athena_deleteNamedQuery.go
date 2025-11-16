package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_deleteNamedQueryCmd = &cobra.Command{
	Use:   "delete-named-query",
	Short: "Deletes the named query if you have access to the workgroup in which the query was saved.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_deleteNamedQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_deleteNamedQueryCmd).Standalone()

		athena_deleteNamedQueryCmd.Flags().String("named-query-id", "", "The unique ID of the query to delete.")
		athena_deleteNamedQueryCmd.MarkFlagRequired("named-query-id")
	})
	athenaCmd.AddCommand(athena_deleteNamedQueryCmd)
}
