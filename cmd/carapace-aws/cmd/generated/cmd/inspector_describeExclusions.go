package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeExclusionsCmd = &cobra.Command{
	Use:   "describe-exclusions",
	Short: "Describes the exclusions that are specified by the exclusions' ARNs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeExclusionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_describeExclusionsCmd).Standalone()

		inspector_describeExclusionsCmd.Flags().String("exclusion-arns", "", "The list of ARNs that specify the exclusions that you want to describe.")
		inspector_describeExclusionsCmd.Flags().String("locale", "", "The locale into which you want to translate the exclusion's title, description, and recommendation.")
		inspector_describeExclusionsCmd.MarkFlagRequired("exclusion-arns")
	})
	inspectorCmd.AddCommand(inspector_describeExclusionsCmd)
}
