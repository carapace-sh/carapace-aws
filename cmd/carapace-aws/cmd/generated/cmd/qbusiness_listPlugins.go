package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listPluginsCmd = &cobra.Command{
	Use:   "list-plugins",
	Short: "Lists configured Amazon Q Business plugins.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listPluginsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listPluginsCmd).Standalone()

		qbusiness_listPluginsCmd.Flags().String("application-id", "", "The identifier of the application the plugin is attached to.")
		qbusiness_listPluginsCmd.Flags().String("max-results", "", "The maximum number of documents to return.")
		qbusiness_listPluginsCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
		qbusiness_listPluginsCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_listPluginsCmd)
}
