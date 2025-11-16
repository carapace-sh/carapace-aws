package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getExecutionPreviewCmd = &cobra.Command{
	Use:   "get-execution-preview",
	Short: "Initiates the process of retrieving an existing preview that shows the effects that running a specified Automation runbook would have on the targeted resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getExecutionPreviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getExecutionPreviewCmd).Standalone()

		ssm_getExecutionPreviewCmd.Flags().String("execution-preview-id", "", "The ID of the existing execution preview.")
		ssm_getExecutionPreviewCmd.MarkFlagRequired("execution-preview-id")
	})
	ssmCmd.AddCommand(ssm_getExecutionPreviewCmd)
}
