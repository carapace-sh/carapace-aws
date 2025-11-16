package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeSupportedLanguagesCmd = &cobra.Command{
	Use:   "describe-supported-languages",
	Short: "Returns a list of supported languages for a specified `categoryCode`, `issueType` and `serviceCode`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeSupportedLanguagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(support_describeSupportedLanguagesCmd).Standalone()

		support_describeSupportedLanguagesCmd.Flags().String("category-code", "", "The category of problem for the support case.")
		support_describeSupportedLanguagesCmd.Flags().String("issue-type", "", "The type of issue for the case.")
		support_describeSupportedLanguagesCmd.Flags().String("service-code", "", "The code for the Amazon Web Services service.")
		support_describeSupportedLanguagesCmd.MarkFlagRequired("category-code")
		support_describeSupportedLanguagesCmd.MarkFlagRequired("issue-type")
		support_describeSupportedLanguagesCmd.MarkFlagRequired("service-code")
	})
	supportCmd.AddCommand(support_describeSupportedLanguagesCmd)
}
