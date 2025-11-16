package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_untagResourceCmd).Standalone()

	launchWizard_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	launchWizard_untagResourceCmd.Flags().String("tag-keys", "", "Keys identifying the tags to remove.")
	launchWizard_untagResourceCmd.MarkFlagRequired("resource-arn")
	launchWizard_untagResourceCmd.MarkFlagRequired("tag-keys")
	launchWizardCmd.AddCommand(launchWizard_untagResourceCmd)
}
