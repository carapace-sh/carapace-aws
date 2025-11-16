package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchUpdateMemberEc2DeepInspectionStatusCmd = &cobra.Command{
	Use:   "batch-update-member-ec2-deep-inspection-status",
	Short: "Activates or deactivates Amazon Inspector deep inspection for the provided member accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchUpdateMemberEc2DeepInspectionStatusCmd).Standalone()

	inspector2_batchUpdateMemberEc2DeepInspectionStatusCmd.Flags().String("account-ids", "", "The unique identifiers for the Amazon Web Services accounts to change Amazon Inspector deep inspection status for.")
	inspector2_batchUpdateMemberEc2DeepInspectionStatusCmd.MarkFlagRequired("account-ids")
	inspector2Cmd.AddCommand(inspector2_batchUpdateMemberEc2DeepInspectionStatusCmd)
}
