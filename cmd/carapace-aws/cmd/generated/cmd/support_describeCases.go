package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeCasesCmd = &cobra.Command{
	Use:   "describe-cases",
	Short: "Returns a list of cases that you specify by passing one or more case IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeCasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(support_describeCasesCmd).Standalone()

		support_describeCasesCmd.Flags().String("after-time", "", "The start date for a filtered date search on support case communications.")
		support_describeCasesCmd.Flags().String("before-time", "", "The end date for a filtered date search on support case communications.")
		support_describeCasesCmd.Flags().String("case-id-list", "", "A list of ID numbers of the support cases you want returned.")
		support_describeCasesCmd.Flags().String("display-id", "", "The ID displayed for a case in the Amazon Web Services Support Center user interface.")
		support_describeCasesCmd.Flags().String("include-communications", "", "Specifies whether to include communications in the `DescribeCases` response.")
		support_describeCasesCmd.Flags().String("include-resolved-cases", "", "Specifies whether to include resolved support cases in the `DescribeCases` response.")
		support_describeCasesCmd.Flags().String("language", "", "The language in which Amazon Web Services Support handles the case.")
		support_describeCasesCmd.Flags().String("max-results", "", "The maximum number of results to return before paginating.")
		support_describeCasesCmd.Flags().String("next-token", "", "A resumption point for pagination.")
	})
	supportCmd.AddCommand(support_describeCasesCmd)
}
