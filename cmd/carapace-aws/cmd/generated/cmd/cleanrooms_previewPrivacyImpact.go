package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_previewPrivacyImpactCmd = &cobra.Command{
	Use:   "preview-privacy-impact",
	Short: "An estimate of the number of aggregation functions that the member who can query can run given epsilon and noise parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_previewPrivacyImpactCmd).Standalone()

	cleanrooms_previewPrivacyImpactCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
	cleanrooms_previewPrivacyImpactCmd.Flags().String("parameters", "", "Specifies the desired epsilon and noise parameters to preview.")
	cleanrooms_previewPrivacyImpactCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_previewPrivacyImpactCmd.MarkFlagRequired("parameters")
	cleanroomsCmd.AddCommand(cleanrooms_previewPrivacyImpactCmd)
}
