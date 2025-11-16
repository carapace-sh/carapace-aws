package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_updateFindingsCmd = &cobra.Command{
	Use:   "update-findings",
	Short: "Updates the status for the specified findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_updateFindingsCmd).Standalone()

	accessanalyzer_updateFindingsCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) that generated the findings to update.")
	accessanalyzer_updateFindingsCmd.Flags().String("client-token", "", "A client token.")
	accessanalyzer_updateFindingsCmd.Flags().String("ids", "", "The IDs of the findings to update.")
	accessanalyzer_updateFindingsCmd.Flags().String("resource-arn", "", "The ARN of the resource identified in the finding.")
	accessanalyzer_updateFindingsCmd.Flags().String("status", "", "The state represents the action to take to update the finding Status.")
	accessanalyzer_updateFindingsCmd.MarkFlagRequired("analyzer-arn")
	accessanalyzer_updateFindingsCmd.MarkFlagRequired("status")
	accessanalyzerCmd.AddCommand(accessanalyzer_updateFindingsCmd)
}
