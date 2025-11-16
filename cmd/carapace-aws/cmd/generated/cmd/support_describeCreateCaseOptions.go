package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeCreateCaseOptionsCmd = &cobra.Command{
	Use:   "describe-create-case-options",
	Short: "Returns a list of CreateCaseOption types along with the corresponding supported hours and language availability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeCreateCaseOptionsCmd).Standalone()

	support_describeCreateCaseOptionsCmd.Flags().String("category-code", "", "The category of problem for the support case.")
	support_describeCreateCaseOptionsCmd.Flags().String("issue-type", "", "The type of issue for the case.")
	support_describeCreateCaseOptionsCmd.Flags().String("language", "", "The language in which Amazon Web Services Support handles the case.")
	support_describeCreateCaseOptionsCmd.Flags().String("service-code", "", "The code for the Amazon Web Services service.")
	support_describeCreateCaseOptionsCmd.MarkFlagRequired("category-code")
	support_describeCreateCaseOptionsCmd.MarkFlagRequired("issue-type")
	support_describeCreateCaseOptionsCmd.MarkFlagRequired("language")
	support_describeCreateCaseOptionsCmd.MarkFlagRequired("service-code")
	supportCmd.AddCommand(support_describeCreateCaseOptionsCmd)
}
