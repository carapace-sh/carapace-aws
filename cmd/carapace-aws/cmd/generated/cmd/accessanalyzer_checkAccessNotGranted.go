package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_checkAccessNotGrantedCmd = &cobra.Command{
	Use:   "check-access-not-granted",
	Short: "Checks whether the specified access isn't allowed by a policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_checkAccessNotGrantedCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_checkAccessNotGrantedCmd).Standalone()

		accessanalyzer_checkAccessNotGrantedCmd.Flags().String("access", "", "An access object containing the permissions that shouldn't be granted by the specified policy.")
		accessanalyzer_checkAccessNotGrantedCmd.Flags().String("policy-document", "", "The JSON policy document to use as the content for the policy.")
		accessanalyzer_checkAccessNotGrantedCmd.Flags().String("policy-type", "", "The type of policy.")
		accessanalyzer_checkAccessNotGrantedCmd.MarkFlagRequired("access")
		accessanalyzer_checkAccessNotGrantedCmd.MarkFlagRequired("policy-document")
		accessanalyzer_checkAccessNotGrantedCmd.MarkFlagRequired("policy-type")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_checkAccessNotGrantedCmd)
}
