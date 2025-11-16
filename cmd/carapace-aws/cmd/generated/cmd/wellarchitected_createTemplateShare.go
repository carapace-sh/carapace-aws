package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createTemplateShareCmd = &cobra.Command{
	Use:   "create-template-share",
	Short: "Create a review template share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createTemplateShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_createTemplateShareCmd).Standalone()

		wellarchitected_createTemplateShareCmd.Flags().String("client-request-token", "", "")
		wellarchitected_createTemplateShareCmd.Flags().String("shared-with", "", "")
		wellarchitected_createTemplateShareCmd.Flags().String("template-arn", "", "The review template ARN.")
		wellarchitected_createTemplateShareCmd.MarkFlagRequired("client-request-token")
		wellarchitected_createTemplateShareCmd.MarkFlagRequired("shared-with")
		wellarchitected_createTemplateShareCmd.MarkFlagRequired("template-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_createTemplateShareCmd)
}
