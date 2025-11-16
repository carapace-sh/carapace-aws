package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_checkNoNewAccessCmd = &cobra.Command{
	Use:   "check-no-new-access",
	Short: "Checks whether new access is allowed for an updated policy when compared to the existing policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_checkNoNewAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_checkNoNewAccessCmd).Standalone()

		accessanalyzer_checkNoNewAccessCmd.Flags().String("existing-policy-document", "", "The JSON policy document to use as the content for the existing policy.")
		accessanalyzer_checkNoNewAccessCmd.Flags().String("new-policy-document", "", "The JSON policy document to use as the content for the updated policy.")
		accessanalyzer_checkNoNewAccessCmd.Flags().String("policy-type", "", "The type of policy to compare.")
		accessanalyzer_checkNoNewAccessCmd.MarkFlagRequired("existing-policy-document")
		accessanalyzer_checkNoNewAccessCmd.MarkFlagRequired("new-policy-document")
		accessanalyzer_checkNoNewAccessCmd.MarkFlagRequired("policy-type")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_checkNoNewAccessCmd)
}
