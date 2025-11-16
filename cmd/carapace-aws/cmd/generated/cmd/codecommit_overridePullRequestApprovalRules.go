package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_overridePullRequestApprovalRulesCmd = &cobra.Command{
	Use:   "override-pull-request-approval-rules",
	Short: "Sets aside (overrides) all approval rule requirements for a specified pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_overridePullRequestApprovalRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_overridePullRequestApprovalRulesCmd).Standalone()

		codecommit_overridePullRequestApprovalRulesCmd.Flags().String("override-status", "", "Whether you want to set aside approval rule requirements for the pull request (OVERRIDE) or revoke a previous override and apply approval rule requirements (REVOKE).")
		codecommit_overridePullRequestApprovalRulesCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request for which you want to override all approval rule requirements.")
		codecommit_overridePullRequestApprovalRulesCmd.Flags().String("revision-id", "", "The system-generated ID of the most recent revision of the pull request.")
		codecommit_overridePullRequestApprovalRulesCmd.MarkFlagRequired("override-status")
		codecommit_overridePullRequestApprovalRulesCmd.MarkFlagRequired("pull-request-id")
		codecommit_overridePullRequestApprovalRulesCmd.MarkFlagRequired("revision-id")
	})
	codecommitCmd.AddCommand(codecommit_overridePullRequestApprovalRulesCmd)
}
