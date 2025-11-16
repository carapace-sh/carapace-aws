package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppAssessmentsCmd = &cobra.Command{
	Use:   "list-app-assessments",
	Short: "Lists the assessments for an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppAssessmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listAppAssessmentsCmd).Standalone()

		resiliencehub_listAppAssessmentsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_listAppAssessmentsCmd.Flags().String("assessment-name", "", "The name for the assessment.")
		resiliencehub_listAppAssessmentsCmd.Flags().String("assessment-status", "", "The current status of the assessment for the resiliency policy.")
		resiliencehub_listAppAssessmentsCmd.Flags().String("compliance-status", "", "The current status of compliance for the resiliency policy.")
		resiliencehub_listAppAssessmentsCmd.Flags().String("invoker", "", "Specifies the entity that invoked a specific assessment, either a `User` or the `System`.")
		resiliencehub_listAppAssessmentsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
		resiliencehub_listAppAssessmentsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listAppAssessmentsCmd.Flags().String("reverse-order", "", "The default is to sort by ascending **startTime**.")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listAppAssessmentsCmd)
}
