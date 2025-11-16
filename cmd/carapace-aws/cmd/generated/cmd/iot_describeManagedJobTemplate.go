package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeManagedJobTemplateCmd = &cobra.Command{
	Use:   "describe-managed-job-template",
	Short: "View details of a managed job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeManagedJobTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeManagedJobTemplateCmd).Standalone()

		iot_describeManagedJobTemplateCmd.Flags().String("template-name", "", "The unique name of a managed job template, which is required.")
		iot_describeManagedJobTemplateCmd.Flags().String("template-version", "", "An optional parameter to specify version of a managed template.")
		iot_describeManagedJobTemplateCmd.MarkFlagRequired("template-name")
	})
	iotCmd.AddCommand(iot_describeManagedJobTemplateCmd)
}
