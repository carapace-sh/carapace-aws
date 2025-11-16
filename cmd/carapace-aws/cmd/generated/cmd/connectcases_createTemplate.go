package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Creates a template in the Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_createTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_createTemplateCmd).Standalone()

		connectcases_createTemplateCmd.Flags().String("description", "", "A brief description of the template.")
		connectcases_createTemplateCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_createTemplateCmd.Flags().String("layout-configuration", "", "Configuration of layouts associated to the template.")
		connectcases_createTemplateCmd.Flags().String("name", "", "A name for the template.")
		connectcases_createTemplateCmd.Flags().String("required-fields", "", "A list of fields that must contain a value for a case to be successfully created with this template.")
		connectcases_createTemplateCmd.Flags().String("rules", "", "A list of case rules (also known as [case field conditions](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html)) on a template.")
		connectcases_createTemplateCmd.Flags().String("status", "", "The status of the template.")
		connectcases_createTemplateCmd.MarkFlagRequired("domain-id")
		connectcases_createTemplateCmd.MarkFlagRequired("name")
	})
	connectcasesCmd.AddCommand(connectcases_createTemplateCmd)
}
