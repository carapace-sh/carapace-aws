package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_createExclusionsPreviewCmd = &cobra.Command{
	Use:   "create-exclusions-preview",
	Short: "Starts the generation of an exclusions preview for the specified assessment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_createExclusionsPreviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_createExclusionsPreviewCmd).Standalone()

		inspector_createExclusionsPreviewCmd.Flags().String("assessment-template-arn", "", "The ARN that specifies the assessment template for which you want to create an exclusions preview.")
		inspector_createExclusionsPreviewCmd.MarkFlagRequired("assessment-template-arn")
	})
	inspectorCmd.AddCommand(inspector_createExclusionsPreviewCmd)
}
