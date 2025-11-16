package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_checkNoPublicAccessCmd = &cobra.Command{
	Use:   "check-no-public-access",
	Short: "Checks whether a resource policy can grant public access to the specified resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_checkNoPublicAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_checkNoPublicAccessCmd).Standalone()

		accessanalyzer_checkNoPublicAccessCmd.Flags().String("policy-document", "", "The JSON policy document to evaluate for public access.")
		accessanalyzer_checkNoPublicAccessCmd.Flags().String("resource-type", "", "The type of resource to evaluate for public access.")
		accessanalyzer_checkNoPublicAccessCmd.MarkFlagRequired("policy-document")
		accessanalyzer_checkNoPublicAccessCmd.MarkFlagRequired("resource-type")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_checkNoPublicAccessCmd)
}
