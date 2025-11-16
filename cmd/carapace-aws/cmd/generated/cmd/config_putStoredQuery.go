package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putStoredQueryCmd = &cobra.Command{
	Use:   "put-stored-query",
	Short: "Saves a new query or updates an existing saved query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putStoredQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putStoredQueryCmd).Standalone()

		config_putStoredQueryCmd.Flags().String("stored-query", "", "A list of `StoredQuery` objects.")
		config_putStoredQueryCmd.Flags().String("tags", "", "A list of `Tags` object.")
		config_putStoredQueryCmd.MarkFlagRequired("stored-query")
	})
	configCmd.AddCommand(config_putStoredQueryCmd)
}
