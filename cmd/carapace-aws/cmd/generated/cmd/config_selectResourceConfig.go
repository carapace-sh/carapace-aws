package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_selectResourceConfigCmd = &cobra.Command{
	Use:   "select-resource-config",
	Short: "Accepts a structured query language (SQL) `SELECT` command, performs the corresponding search, and returns resource configurations matching the properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_selectResourceConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_selectResourceConfigCmd).Standalone()

		config_selectResourceConfigCmd.Flags().String("expression", "", "The SQL query `SELECT` command.")
		config_selectResourceConfigCmd.Flags().String("limit", "", "The maximum number of query results returned on each page.")
		config_selectResourceConfigCmd.Flags().String("next-token", "", "The `nextToken` string returned in a previous request that you use to request the next page of results in a paginated response.")
		config_selectResourceConfigCmd.MarkFlagRequired("expression")
	})
	configCmd.AddCommand(config_selectResourceConfigCmd)
}
