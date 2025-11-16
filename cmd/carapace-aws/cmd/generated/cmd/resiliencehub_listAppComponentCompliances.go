package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppComponentCompliancesCmd = &cobra.Command{
	Use:   "list-app-component-compliances",
	Short: "Lists the compliances for an Resilience Hub Application Component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppComponentCompliancesCmd).Standalone()

	resiliencehub_listAppComponentCompliancesCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
	resiliencehub_listAppComponentCompliancesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listAppComponentCompliancesCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listAppComponentCompliancesCmd.MarkFlagRequired("assessment-arn")
	resiliencehubCmd.AddCommand(resiliencehub_listAppComponentCompliancesCmd)
}
