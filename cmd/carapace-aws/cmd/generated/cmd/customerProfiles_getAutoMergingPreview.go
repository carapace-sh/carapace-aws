package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getAutoMergingPreviewCmd = &cobra.Command{
	Use:   "get-auto-merging-preview",
	Short: "Tests the auto-merging settings of your Identity Resolution Job without merging your data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getAutoMergingPreviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getAutoMergingPreviewCmd).Standalone()

		customerProfiles_getAutoMergingPreviewCmd.Flags().String("conflict-resolution", "", "How the auto-merging process should resolve conflicts between different profiles.")
		customerProfiles_getAutoMergingPreviewCmd.Flags().String("consolidation", "", "A list of matching attributes that represent matching criteria.")
		customerProfiles_getAutoMergingPreviewCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getAutoMergingPreviewCmd.Flags().String("min-allowed-confidence-score-for-merging", "", "Minimum confidence score required for profiles within a matching group to be merged during the auto-merge process.")
		customerProfiles_getAutoMergingPreviewCmd.MarkFlagRequired("conflict-resolution")
		customerProfiles_getAutoMergingPreviewCmd.MarkFlagRequired("consolidation")
		customerProfiles_getAutoMergingPreviewCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getAutoMergingPreviewCmd)
}
