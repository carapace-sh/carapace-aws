package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteStoredQueryCmd = &cobra.Command{
	Use:   "delete-stored-query",
	Short: "Deletes the stored query for a single Amazon Web Services account and a single Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteStoredQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteStoredQueryCmd).Standalone()

		config_deleteStoredQueryCmd.Flags().String("query-name", "", "The name of the query that you want to delete.")
		config_deleteStoredQueryCmd.MarkFlagRequired("query-name")
	})
	configCmd.AddCommand(config_deleteStoredQueryCmd)
}
