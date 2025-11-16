package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_cancelPolicyGenerationCmd = &cobra.Command{
	Use:   "cancel-policy-generation",
	Short: "Cancels the requested policy generation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_cancelPolicyGenerationCmd).Standalone()

	accessanalyzer_cancelPolicyGenerationCmd.Flags().String("job-id", "", "The `JobId` that is returned by the `StartPolicyGeneration` operation.")
	accessanalyzer_cancelPolicyGenerationCmd.MarkFlagRequired("job-id")
	accessanalyzerCmd.AddCommand(accessanalyzer_cancelPolicyGenerationCmd)
}
