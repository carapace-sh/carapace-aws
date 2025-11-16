package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getMatchIdCmd = &cobra.Command{
	Use:   "get-match-id",
	Short: "Returns the corresponding Match ID of a customer record if the record has been processed in a rule-based matching workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getMatchIdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_getMatchIdCmd).Standalone()

		entityresolution_getMatchIdCmd.Flags().Bool("apply-normalization", false, "Normalizes the attributes defined in the schema in the input data.")
		entityresolution_getMatchIdCmd.Flags().Bool("no-apply-normalization", false, "Normalizes the attributes defined in the schema in the input data.")
		entityresolution_getMatchIdCmd.Flags().String("record", "", "The record to fetch the Match ID for.")
		entityresolution_getMatchIdCmd.Flags().String("workflow-name", "", "The name of the workflow.")
		entityresolution_getMatchIdCmd.Flag("no-apply-normalization").Hidden = true
		entityresolution_getMatchIdCmd.MarkFlagRequired("record")
		entityresolution_getMatchIdCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_getMatchIdCmd)
}
