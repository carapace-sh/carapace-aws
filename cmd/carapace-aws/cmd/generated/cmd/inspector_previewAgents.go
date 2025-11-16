package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_previewAgentsCmd = &cobra.Command{
	Use:   "preview-agents",
	Short: "Previews the agents installed on the EC2 instances that are part of the specified assessment target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_previewAgentsCmd).Standalone()

	inspector_previewAgentsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
	inspector_previewAgentsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	inspector_previewAgentsCmd.Flags().String("preview-agents-arn", "", "The ARN of the assessment target whose agents you want to preview.")
	inspector_previewAgentsCmd.MarkFlagRequired("preview-agents-arn")
	inspectorCmd.AddCommand(inspector_previewAgentsCmd)
}
