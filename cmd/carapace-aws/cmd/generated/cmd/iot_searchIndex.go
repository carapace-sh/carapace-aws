package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_searchIndexCmd = &cobra.Command{
	Use:   "search-index",
	Short: "The query search index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_searchIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_searchIndexCmd).Standalone()

		iot_searchIndexCmd.Flags().String("index-name", "", "The search index name.")
		iot_searchIndexCmd.Flags().String("max-results", "", "The maximum number of results to return per page at one time.")
		iot_searchIndexCmd.Flags().String("next-token", "", "The token used to get the next set of results, or `null` if there are no additional results.")
		iot_searchIndexCmd.Flags().String("query-string", "", "The search query string.")
		iot_searchIndexCmd.Flags().String("query-version", "", "The query version.")
		iot_searchIndexCmd.MarkFlagRequired("query-string")
	})
	iotCmd.AddCommand(iot_searchIndexCmd)
}
