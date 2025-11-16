package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listEndpointsCmd = &cobra.Command{
	Use:   "list-endpoints",
	Short: "Gets a list of all existing endpoints that you've created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listEndpointsCmd).Standalone()

	comprehend_listEndpointsCmd.Flags().String("filter", "", "Filters the endpoints that are returned.")
	comprehend_listEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehend_listEndpointsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listEndpointsCmd)
}
