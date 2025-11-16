package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_tagResourceCmd).Standalone()

	launchWizard_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	launchWizard_tagResourceCmd.Flags().String("tags", "", "One or more tags to attach to the resource.")
	launchWizard_tagResourceCmd.MarkFlagRequired("resource-arn")
	launchWizard_tagResourceCmd.MarkFlagRequired("tags")
	launchWizardCmd.AddCommand(launchWizard_tagResourceCmd)
}
