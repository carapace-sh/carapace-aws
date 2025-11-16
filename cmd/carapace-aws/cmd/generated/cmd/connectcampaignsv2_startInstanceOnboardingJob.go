package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_startInstanceOnboardingJobCmd = &cobra.Command{
	Use:   "start-instance-onboarding-job",
	Short: "Onboard the specific Amazon Connect instance to Connect Campaigns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_startInstanceOnboardingJobCmd).Standalone()

	connectcampaignsv2_startInstanceOnboardingJobCmd.Flags().String("connect-instance-id", "", "")
	connectcampaignsv2_startInstanceOnboardingJobCmd.Flags().String("encryption-config", "", "")
	connectcampaignsv2_startInstanceOnboardingJobCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsv2_startInstanceOnboardingJobCmd.MarkFlagRequired("encryption-config")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_startInstanceOnboardingJobCmd)
}
