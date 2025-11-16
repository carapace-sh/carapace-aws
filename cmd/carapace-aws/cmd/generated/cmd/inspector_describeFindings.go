package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeFindingsCmd = &cobra.Command{
	Use:   "describe-findings",
	Short: "Describes the findings that are specified by the ARNs of the findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_describeFindingsCmd).Standalone()

		inspector_describeFindingsCmd.Flags().String("finding-arns", "", "The ARN that specifies the finding that you want to describe.")
		inspector_describeFindingsCmd.Flags().String("locale", "", "The locale into which you want to translate a finding description, recommendation, and the short description that identifies the finding.")
		inspector_describeFindingsCmd.MarkFlagRequired("finding-arns")
	})
	inspectorCmd.AddCommand(inspector_describeFindingsCmd)
}
