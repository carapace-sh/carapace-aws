package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchGetMemberEc2DeepInspectionStatusCmd = &cobra.Command{
	Use:   "batch-get-member-ec2-deep-inspection-status",
	Short: "Retrieves Amazon Inspector deep inspection activation status of multiple member accounts within your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchGetMemberEc2DeepInspectionStatusCmd).Standalone()

	inspector2_batchGetMemberEc2DeepInspectionStatusCmd.Flags().String("account-ids", "", "The unique identifiers for the Amazon Web Services accounts to retrieve Amazon Inspector deep inspection activation status for.")
	inspector2Cmd.AddCommand(inspector2_batchGetMemberEc2DeepInspectionStatusCmd)
}
