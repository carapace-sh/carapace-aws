package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a response plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_tagResourceCmd).Standalone()

	ssmIncidents_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the response plan you're adding the tags to.")
	ssmIncidents_tagResourceCmd.Flags().String("tags", "", "A list of tags to add to the response plan.")
	ssmIncidents_tagResourceCmd.MarkFlagRequired("resource-arn")
	ssmIncidents_tagResourceCmd.MarkFlagRequired("tags")
	ssmIncidentsCmd.AddCommand(ssmIncidents_tagResourceCmd)
}
