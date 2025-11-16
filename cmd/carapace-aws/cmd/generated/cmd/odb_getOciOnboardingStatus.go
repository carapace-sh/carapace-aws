package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getOciOnboardingStatusCmd = &cobra.Command{
	Use:   "get-oci-onboarding-status",
	Short: "Returns the tenancy activation link and onboarding status for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getOciOnboardingStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_getOciOnboardingStatusCmd).Standalone()

	})
	odbCmd.AddCommand(odb_getOciOnboardingStatusCmd)
}
