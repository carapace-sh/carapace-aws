package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getGeneratedPolicyCmd = &cobra.Command{
	Use:   "get-generated-policy",
	Short: "Retrieves the policy that was generated using `StartPolicyGeneration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getGeneratedPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_getGeneratedPolicyCmd).Standalone()

		accessanalyzer_getGeneratedPolicyCmd.Flags().Bool("include-resource-placeholders", false, "The level of detail that you want to generate.")
		accessanalyzer_getGeneratedPolicyCmd.Flags().Bool("include-service-level-template", false, "The level of detail that you want to generate.")
		accessanalyzer_getGeneratedPolicyCmd.Flags().String("job-id", "", "The `JobId` that is returned by the `StartPolicyGeneration` operation.")
		accessanalyzer_getGeneratedPolicyCmd.Flags().Bool("no-include-resource-placeholders", false, "The level of detail that you want to generate.")
		accessanalyzer_getGeneratedPolicyCmd.Flags().Bool("no-include-service-level-template", false, "The level of detail that you want to generate.")
		accessanalyzer_getGeneratedPolicyCmd.MarkFlagRequired("job-id")
		accessanalyzer_getGeneratedPolicyCmd.Flag("no-include-resource-placeholders").Hidden = true
		accessanalyzer_getGeneratedPolicyCmd.Flag("no-include-service-level-template").Hidden = true
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_getGeneratedPolicyCmd)
}
