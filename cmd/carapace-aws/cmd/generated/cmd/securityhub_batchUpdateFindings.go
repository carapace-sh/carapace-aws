package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchUpdateFindingsCmd = &cobra.Command{
	Use:   "batch-update-findings",
	Short: "Used by Security Hub customers to update information about their investigation into one or more findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchUpdateFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchUpdateFindingsCmd).Standalone()

		securityhub_batchUpdateFindingsCmd.Flags().String("confidence", "", "The updated value for the finding confidence.")
		securityhub_batchUpdateFindingsCmd.Flags().String("criticality", "", "The updated value for the level of importance assigned to the resources associated with the findings.")
		securityhub_batchUpdateFindingsCmd.Flags().String("finding-identifiers", "", "The list of findings to update.")
		securityhub_batchUpdateFindingsCmd.Flags().String("note", "", "")
		securityhub_batchUpdateFindingsCmd.Flags().String("related-findings", "", "A list of findings that are related to the updated findings.")
		securityhub_batchUpdateFindingsCmd.Flags().String("severity", "", "Used to update the finding severity.")
		securityhub_batchUpdateFindingsCmd.Flags().String("types", "", "One or more finding types in the format of namespace/category/classifier that classify a finding.")
		securityhub_batchUpdateFindingsCmd.Flags().String("user-defined-fields", "", "A list of name/value string pairs associated with the finding.")
		securityhub_batchUpdateFindingsCmd.Flags().String("verification-state", "", "Indicates the veracity of a finding.")
		securityhub_batchUpdateFindingsCmd.Flags().String("workflow", "", "Used to update the workflow status of a finding.")
		securityhub_batchUpdateFindingsCmd.MarkFlagRequired("finding-identifiers")
	})
	securityhubCmd.AddCommand(securityhub_batchUpdateFindingsCmd)
}
