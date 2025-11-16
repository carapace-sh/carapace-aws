package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getStoredQueryCmd = &cobra.Command{
	Use:   "get-stored-query",
	Short: "Returns the details of a specific stored query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getStoredQueryCmd).Standalone()

	config_getStoredQueryCmd.Flags().String("query-name", "", "The name of the query.")
	config_getStoredQueryCmd.MarkFlagRequired("query-name")
	configCmd.AddCommand(config_getStoredQueryCmd)
}
