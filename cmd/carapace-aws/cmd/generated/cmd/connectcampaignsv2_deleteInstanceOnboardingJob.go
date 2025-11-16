package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_deleteInstanceOnboardingJobCmd = &cobra.Command{
	Use:   "delete-instance-onboarding-job",
	Short: "Delete the Connect Campaigns onboarding job for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_deleteInstanceOnboardingJobCmd).Standalone()

	connectcampaignsv2_deleteInstanceOnboardingJobCmd.Flags().String("connect-instance-id", "", "")
	connectcampaignsv2_deleteInstanceOnboardingJobCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_deleteInstanceOnboardingJobCmd)
}
