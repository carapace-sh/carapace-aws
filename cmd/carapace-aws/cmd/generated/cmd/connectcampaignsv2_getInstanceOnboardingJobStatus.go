package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_getInstanceOnboardingJobStatusCmd = &cobra.Command{
	Use:   "get-instance-onboarding-job-status",
	Short: "Get the specific instance onboarding job status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_getInstanceOnboardingJobStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_getInstanceOnboardingJobStatusCmd).Standalone()

		connectcampaignsv2_getInstanceOnboardingJobStatusCmd.Flags().String("connect-instance-id", "", "")
		connectcampaignsv2_getInstanceOnboardingJobStatusCmd.MarkFlagRequired("connect-instance-id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_getInstanceOnboardingJobStatusCmd)
}
