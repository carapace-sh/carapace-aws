package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getDataSourceIntrospectionCmd = &cobra.Command{
	Use:   "get-data-source-introspection",
	Short: "Retrieves the record of an existing introspection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getDataSourceIntrospectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getDataSourceIntrospectionCmd).Standalone()

		appsync_getDataSourceIntrospectionCmd.Flags().Bool("include-models-sdl", false, "A boolean flag that determines whether SDL should be generated for introspected types.")
		appsync_getDataSourceIntrospectionCmd.Flags().String("introspection-id", "", "The introspection ID.")
		appsync_getDataSourceIntrospectionCmd.Flags().String("max-results", "", "The maximum number of introspected types that will be returned in a single response.")
		appsync_getDataSourceIntrospectionCmd.Flags().String("next-token", "", "Determines the number of types to be returned in a single response before paginating.")
		appsync_getDataSourceIntrospectionCmd.Flags().Bool("no-include-models-sdl", false, "A boolean flag that determines whether SDL should be generated for introspected types.")
		appsync_getDataSourceIntrospectionCmd.MarkFlagRequired("introspection-id")
		appsync_getDataSourceIntrospectionCmd.Flag("no-include-models-sdl").Hidden = true
	})
	appsyncCmd.AddCommand(appsync_getDataSourceIntrospectionCmd)
}
