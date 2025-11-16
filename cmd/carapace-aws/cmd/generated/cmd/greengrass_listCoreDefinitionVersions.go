package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listCoreDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-core-definition-versions",
	Short: "Lists the versions of a core definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listCoreDefinitionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listCoreDefinitionVersionsCmd).Standalone()

		greengrass_listCoreDefinitionVersionsCmd.Flags().String("core-definition-id", "", "The ID of the core definition.")
		greengrass_listCoreDefinitionVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listCoreDefinitionVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_listCoreDefinitionVersionsCmd.MarkFlagRequired("core-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_listCoreDefinitionVersionsCmd)
}
