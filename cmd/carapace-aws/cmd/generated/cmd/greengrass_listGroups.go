package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Retrieves a list of groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listGroupsCmd).Standalone()

		greengrass_listGroupsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listGroupsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	})
	greengrassCmd.AddCommand(greengrass_listGroupsCmd)
}
