package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getSystemTemplateRevisionsCmd = &cobra.Command{
	Use:   "get-system-template-revisions",
	Short: "Gets revisions made to the specified system template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getSystemTemplateRevisionsCmd).Standalone()

	iotthingsgraph_getSystemTemplateRevisionsCmd.Flags().String("id", "", "The ID of the system template.")
	iotthingsgraph_getSystemTemplateRevisionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	iotthingsgraph_getSystemTemplateRevisionsCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iotthingsgraph_getSystemTemplateRevisionsCmd.MarkFlagRequired("id")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_getSystemTemplateRevisionsCmd)
}
