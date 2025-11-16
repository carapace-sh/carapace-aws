package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeQpersonalizationConfigurationCmd = &cobra.Command{
	Use:   "describe-qpersonalization-configuration",
	Short: "Describes a personalization configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeQpersonalizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeQpersonalizationConfigurationCmd).Standalone()

		quicksight_describeQpersonalizationConfigurationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the personalization configuration that the user wants described.")
		quicksight_describeQpersonalizationConfigurationCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_describeQpersonalizationConfigurationCmd)
}
