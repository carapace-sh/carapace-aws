package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_startExecutionPreviewCmd = &cobra.Command{
	Use:   "start-execution-preview",
	Short: "Initiates the process of creating a preview showing the effects that running a specified Automation runbook would have on the targeted resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_startExecutionPreviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_startExecutionPreviewCmd).Standalone()

		ssm_startExecutionPreviewCmd.Flags().String("document-name", "", "The name of the Automation runbook to run.")
		ssm_startExecutionPreviewCmd.Flags().String("document-version", "", "The version of the Automation runbook to run.")
		ssm_startExecutionPreviewCmd.Flags().String("execution-inputs", "", "Information about the inputs that can be specified for the preview operation.")
		ssm_startExecutionPreviewCmd.MarkFlagRequired("document-name")
	})
	ssmCmd.AddCommand(ssm_startExecutionPreviewCmd)
}
