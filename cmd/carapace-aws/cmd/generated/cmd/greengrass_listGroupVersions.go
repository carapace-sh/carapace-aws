package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listGroupVersionsCmd = &cobra.Command{
	Use:   "list-group-versions",
	Short: "Lists the versions of a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listGroupVersionsCmd).Standalone()

	greengrass_listGroupVersionsCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_listGroupVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listGroupVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrass_listGroupVersionsCmd.MarkFlagRequired("group-id")
	greengrassCmd.AddCommand(greengrass_listGroupVersionsCmd)
}
