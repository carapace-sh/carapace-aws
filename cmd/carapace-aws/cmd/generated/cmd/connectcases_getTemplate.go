package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_getTemplateCmd = &cobra.Command{
	Use:   "get-template",
	Short: "Returns the details for the requested template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_getTemplateCmd).Standalone()

	connectcases_getTemplateCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_getTemplateCmd.Flags().String("template-id", "", "A unique identifier of a template.")
	connectcases_getTemplateCmd.MarkFlagRequired("domain-id")
	connectcases_getTemplateCmd.MarkFlagRequired("template-id")
	connectcasesCmd.AddCommand(connectcases_getTemplateCmd)
}
