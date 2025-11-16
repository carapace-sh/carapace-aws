package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_updateTemplateCmd = &cobra.Command{
	Use:   "update-template",
	Short: "Updates the attributes of an existing template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_updateTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_updateTemplateCmd).Standalone()

		connectcases_updateTemplateCmd.Flags().String("description", "", "A brief description of the template.")
		connectcases_updateTemplateCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_updateTemplateCmd.Flags().String("layout-configuration", "", "Configuration of layouts associated to the template.")
		connectcases_updateTemplateCmd.Flags().String("name", "", "The name of the template.")
		connectcases_updateTemplateCmd.Flags().String("required-fields", "", "A list of fields that must contain a value for a case to be successfully created with this template.")
		connectcases_updateTemplateCmd.Flags().String("rules", "", "A list of case rules (also known as [case field conditions](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html)) on a template.")
		connectcases_updateTemplateCmd.Flags().String("status", "", "The status of the template.")
		connectcases_updateTemplateCmd.Flags().String("template-id", "", "A unique identifier for the template.")
		connectcases_updateTemplateCmd.MarkFlagRequired("domain-id")
		connectcases_updateTemplateCmd.MarkFlagRequired("template-id")
	})
	connectcasesCmd.AddCommand(connectcases_updateTemplateCmd)
}
