package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_startInstanceOnboardingJobCmd = &cobra.Command{
	Use:   "start-instance-onboarding-job",
	Short: "Onboard the specific Amazon Connect instance to Connect Campaigns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_startInstanceOnboardingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_startInstanceOnboardingJobCmd).Standalone()

		connectcampaigns_startInstanceOnboardingJobCmd.Flags().String("connect-instance-id", "", "")
		connectcampaigns_startInstanceOnboardingJobCmd.Flags().String("encryption-config", "", "")
		connectcampaigns_startInstanceOnboardingJobCmd.MarkFlagRequired("connect-instance-id")
		connectcampaigns_startInstanceOnboardingJobCmd.MarkFlagRequired("encryption-config")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_startInstanceOnboardingJobCmd)
}
