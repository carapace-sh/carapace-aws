package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_generateMatchIdCmd = &cobra.Command{
	Use:   "generate-match-id",
	Short: "Generates or retrieves Match IDs for records using a rule-based matching workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_generateMatchIdCmd).Standalone()

	entityresolution_generateMatchIdCmd.Flags().String("processing-type", "", "The processing mode that determines how Match IDs are generated and results are saved.")
	entityresolution_generateMatchIdCmd.Flags().String("records", "", "The records to match.")
	entityresolution_generateMatchIdCmd.Flags().String("workflow-name", "", "The name of the rule-based matching workflow.")
	entityresolution_generateMatchIdCmd.MarkFlagRequired("records")
	entityresolution_generateMatchIdCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_generateMatchIdCmd)
}
