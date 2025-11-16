package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listTemplateSharesCmd = &cobra.Command{
	Use:   "list-template-shares",
	Short: "List review template shares.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listTemplateSharesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_listTemplateSharesCmd).Standalone()

		wellarchitected_listTemplateSharesCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
		wellarchitected_listTemplateSharesCmd.Flags().String("next-token", "", "")
		wellarchitected_listTemplateSharesCmd.Flags().String("shared-with-prefix", "", "The Amazon Web Services account ID, organization ID, or organizational unit (OU) ID with which the profile is shared.")
		wellarchitected_listTemplateSharesCmd.Flags().String("status", "", "")
		wellarchitected_listTemplateSharesCmd.Flags().String("template-arn", "", "The review template ARN.")
		wellarchitected_listTemplateSharesCmd.MarkFlagRequired("template-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_listTemplateSharesCmd)
}
