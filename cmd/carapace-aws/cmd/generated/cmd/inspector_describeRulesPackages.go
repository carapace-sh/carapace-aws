package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeRulesPackagesCmd = &cobra.Command{
	Use:   "describe-rules-packages",
	Short: "Describes the rules packages that are specified by the ARNs of the rules packages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeRulesPackagesCmd).Standalone()

	inspector_describeRulesPackagesCmd.Flags().String("locale", "", "The locale that you want to translate a rules package description into.")
	inspector_describeRulesPackagesCmd.Flags().String("rules-package-arns", "", "The ARN that specifies the rules package that you want to describe.")
	inspector_describeRulesPackagesCmd.MarkFlagRequired("rules-package-arns")
	inspectorCmd.AddCommand(inspector_describeRulesPackagesCmd)
}
