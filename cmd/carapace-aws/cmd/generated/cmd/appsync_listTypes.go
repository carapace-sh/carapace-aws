package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listTypesCmd = &cobra.Command{
	Use:   "list-types",
	Short: "Lists the types for a given API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listTypesCmd).Standalone()

	appsync_listTypesCmd.Flags().String("api-id", "", "The API ID.")
	appsync_listTypesCmd.Flags().String("format", "", "The type format: SDL or JSON.")
	appsync_listTypesCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
	appsync_listTypesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
	appsync_listTypesCmd.MarkFlagRequired("api-id")
	appsync_listTypesCmd.MarkFlagRequired("format")
	appsyncCmd.AddCommand(appsync_listTypesCmd)
}
