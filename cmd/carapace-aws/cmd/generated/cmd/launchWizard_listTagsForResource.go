package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags associated with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_listTagsForResourceCmd).Standalone()

		launchWizard_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		launchWizard_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	launchWizardCmd.AddCommand(launchWizard_listTagsForResourceCmd)
}
