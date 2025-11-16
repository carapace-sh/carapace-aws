package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeDraftAppVersionResourcesImportStatusCmd = &cobra.Command{
	Use:   "describe-draft-app-version-resources-import-status",
	Short: "Describes the status of importing resources to an application version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeDraftAppVersionResourcesImportStatusCmd).Standalone()

	resiliencehub_describeDraftAppVersionResourcesImportStatusCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_describeDraftAppVersionResourcesImportStatusCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_describeDraftAppVersionResourcesImportStatusCmd)
}
