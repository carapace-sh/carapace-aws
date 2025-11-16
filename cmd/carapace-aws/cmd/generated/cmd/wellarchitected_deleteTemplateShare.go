package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteTemplateShareCmd = &cobra.Command{
	Use:   "delete-template-share",
	Short: "Delete a review template share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteTemplateShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_deleteTemplateShareCmd).Standalone()

		wellarchitected_deleteTemplateShareCmd.Flags().String("client-request-token", "", "")
		wellarchitected_deleteTemplateShareCmd.Flags().String("share-id", "", "")
		wellarchitected_deleteTemplateShareCmd.Flags().String("template-arn", "", "The review template ARN.")
		wellarchitected_deleteTemplateShareCmd.MarkFlagRequired("client-request-token")
		wellarchitected_deleteTemplateShareCmd.MarkFlagRequired("share-id")
		wellarchitected_deleteTemplateShareCmd.MarkFlagRequired("template-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_deleteTemplateShareCmd)
}
