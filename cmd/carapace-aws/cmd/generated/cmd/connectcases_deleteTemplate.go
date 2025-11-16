package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_deleteTemplateCmd = &cobra.Command{
	Use:   "delete-template",
	Short: "Deletes a cases template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_deleteTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_deleteTemplateCmd).Standalone()

		connectcases_deleteTemplateCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_deleteTemplateCmd.Flags().String("template-id", "", "A unique identifier of a template.")
		connectcases_deleteTemplateCmd.MarkFlagRequired("domain-id")
		connectcases_deleteTemplateCmd.MarkFlagRequired("template-id")
	})
	connectcasesCmd.AddCommand(connectcases_deleteTemplateCmd)
}
