package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_getExclusionsPreviewCmd = &cobra.Command{
	Use:   "get-exclusions-preview",
	Short: "Retrieves the exclusions preview (a list of ExclusionPreview objects) specified by the preview token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_getExclusionsPreviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_getExclusionsPreviewCmd).Standalone()

		inspector_getExclusionsPreviewCmd.Flags().String("assessment-template-arn", "", "The ARN that specifies the assessment template for which the exclusions preview was requested.")
		inspector_getExclusionsPreviewCmd.Flags().String("locale", "", "The locale into which you want to translate the exclusion's title, description, and recommendation.")
		inspector_getExclusionsPreviewCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
		inspector_getExclusionsPreviewCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
		inspector_getExclusionsPreviewCmd.Flags().String("preview-token", "", "The unique identifier associated of the exclusions preview.")
		inspector_getExclusionsPreviewCmd.MarkFlagRequired("assessment-template-arn")
		inspector_getExclusionsPreviewCmd.MarkFlagRequired("preview-token")
	})
	inspectorCmd.AddCommand(inspector_getExclusionsPreviewCmd)
}
