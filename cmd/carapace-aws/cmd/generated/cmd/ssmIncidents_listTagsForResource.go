package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that are attached to the specified response plan or incident.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_listTagsForResourceCmd).Standalone()

		ssmIncidents_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the response plan or incident.")
		ssmIncidents_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_listTagsForResourceCmd)
}
