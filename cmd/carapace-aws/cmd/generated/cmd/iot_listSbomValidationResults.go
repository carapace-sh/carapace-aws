package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listSbomValidationResultsCmd = &cobra.Command{
	Use:   "list-sbom-validation-results",
	Short: "The validation results for all software bill of materials (SBOM) attached to a specific software package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listSbomValidationResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listSbomValidationResultsCmd).Standalone()

		iot_listSbomValidationResultsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listSbomValidationResultsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results, or null if there are no additional results.")
		iot_listSbomValidationResultsCmd.Flags().String("package-name", "", "The name of the new software package.")
		iot_listSbomValidationResultsCmd.Flags().String("validation-result", "", "The end result of the")
		iot_listSbomValidationResultsCmd.Flags().String("version-name", "", "The name of the new package version.")
		iot_listSbomValidationResultsCmd.MarkFlagRequired("package-name")
		iot_listSbomValidationResultsCmd.MarkFlagRequired("version-name")
	})
	iotCmd.AddCommand(iot_listSbomValidationResultsCmd)
}
