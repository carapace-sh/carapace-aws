package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_putDraftAppVersionTemplateCmd = &cobra.Command{
	Use:   "put-draft-app-version-template",
	Short: "Adds or updates the app template for an Resilience Hub application draft version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_putDraftAppVersionTemplateCmd).Standalone()

	resiliencehub_putDraftAppVersionTemplateCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_putDraftAppVersionTemplateCmd.Flags().String("app-template-body", "", "A JSON string that provides information about your application structure.")
	resiliencehub_putDraftAppVersionTemplateCmd.MarkFlagRequired("app-arn")
	resiliencehub_putDraftAppVersionTemplateCmd.MarkFlagRequired("app-template-body")
	resiliencehubCmd.AddCommand(resiliencehub_putDraftAppVersionTemplateCmd)
}
