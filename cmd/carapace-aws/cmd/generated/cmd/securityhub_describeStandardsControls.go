package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeStandardsControlsCmd = &cobra.Command{
	Use:   "describe-standards-controls",
	Short: "Returns a list of security standards controls.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeStandardsControlsCmd).Standalone()

	securityhub_describeStandardsControlsCmd.Flags().String("max-results", "", "The maximum number of security standard controls to return.")
	securityhub_describeStandardsControlsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	securityhub_describeStandardsControlsCmd.Flags().String("standards-subscription-arn", "", "The ARN of a resource that represents your subscription to a supported standard.")
	securityhub_describeStandardsControlsCmd.MarkFlagRequired("standards-subscription-arn")
	securityhubCmd.AddCommand(securityhub_describeStandardsControlsCmd)
}
