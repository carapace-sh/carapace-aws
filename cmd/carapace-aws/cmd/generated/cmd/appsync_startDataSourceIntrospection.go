package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_startDataSourceIntrospectionCmd = &cobra.Command{
	Use:   "start-data-source-introspection",
	Short: "Creates a new introspection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_startDataSourceIntrospectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_startDataSourceIntrospectionCmd).Standalone()

		appsync_startDataSourceIntrospectionCmd.Flags().String("rds-data-api-config", "", "The `rdsDataApiConfig` object data.")
	})
	appsyncCmd.AddCommand(appsync_startDataSourceIntrospectionCmd)
}
