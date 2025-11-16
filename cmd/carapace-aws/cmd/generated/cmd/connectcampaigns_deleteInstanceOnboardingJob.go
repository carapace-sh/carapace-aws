package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_deleteInstanceOnboardingJobCmd = &cobra.Command{
	Use:   "delete-instance-onboarding-job",
	Short: "Delete the Connect Campaigns onboarding job for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_deleteInstanceOnboardingJobCmd).Standalone()

	connectcampaigns_deleteInstanceOnboardingJobCmd.Flags().String("connect-instance-id", "", "")
	connectcampaigns_deleteInstanceOnboardingJobCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsCmd.AddCommand(connectcampaigns_deleteInstanceOnboardingJobCmd)
}
