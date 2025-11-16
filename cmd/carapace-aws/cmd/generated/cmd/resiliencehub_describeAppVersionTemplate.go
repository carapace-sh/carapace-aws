package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeAppVersionTemplateCmd = &cobra.Command{
	Use:   "describe-app-version-template",
	Short: "Describes details about an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeAppVersionTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_describeAppVersionTemplateCmd).Standalone()

		resiliencehub_describeAppVersionTemplateCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_describeAppVersionTemplateCmd.Flags().String("app-version", "", "The version of the application.")
		resiliencehub_describeAppVersionTemplateCmd.MarkFlagRequired("app-arn")
		resiliencehub_describeAppVersionTemplateCmd.MarkFlagRequired("app-version")
	})
	resiliencehubCmd.AddCommand(resiliencehub_describeAppVersionTemplateCmd)
}
