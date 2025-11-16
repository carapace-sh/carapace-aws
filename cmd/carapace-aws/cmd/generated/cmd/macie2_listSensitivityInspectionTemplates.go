package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listSensitivityInspectionTemplatesCmd = &cobra.Command{
	Use:   "list-sensitivity-inspection-templates",
	Short: "Retrieves a subset of information about the sensitivity inspection template for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listSensitivityInspectionTemplatesCmd).Standalone()

	macie2_listSensitivityInspectionTemplatesCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
	macie2_listSensitivityInspectionTemplatesCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2Cmd.AddCommand(macie2_listSensitivityInspectionTemplatesCmd)
}
