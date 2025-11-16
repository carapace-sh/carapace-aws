package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_getInstanceOnboardingJobStatusCmd = &cobra.Command{
	Use:   "get-instance-onboarding-job-status",
	Short: "Get the specific instance onboarding job status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_getInstanceOnboardingJobStatusCmd).Standalone()

	connectcampaigns_getInstanceOnboardingJobStatusCmd.Flags().String("connect-instance-id", "", "")
	connectcampaigns_getInstanceOnboardingJobStatusCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsCmd.AddCommand(connectcampaigns_getInstanceOnboardingJobStatusCmd)
}
