package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_listEnvironmentVpcsCmd = &cobra.Command{
	Use:   "list-environment-vpcs",
	Short: "Lists all Amazon Web Services Migration Hub Refactor Spaces service virtual private clouds (VPCs) that are part of the environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_listEnvironmentVpcsCmd).Standalone()

	migrationHubRefactorSpaces_listEnvironmentVpcsCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
	migrationHubRefactorSpaces_listEnvironmentVpcsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	migrationHubRefactorSpaces_listEnvironmentVpcsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	migrationHubRefactorSpaces_listEnvironmentVpcsCmd.MarkFlagRequired("environment-identifier")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_listEnvironmentVpcsCmd)
}
