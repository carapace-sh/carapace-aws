package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describeReleaseLabelCmd = &cobra.Command{
	Use:   "describe-release-label",
	Short: "Provides Amazon EMR release label details, such as the releases available the Region where the API request is run, and the available applications for a specific Amazon EMR release label.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describeReleaseLabelCmd).Standalone()

	emr_describeReleaseLabelCmd.Flags().String("max-results", "", "Reserved for future use.")
	emr_describeReleaseLabelCmd.Flags().String("next-token", "", "The pagination token.")
	emr_describeReleaseLabelCmd.Flags().String("release-label", "", "The target release label to be described.")
	emrCmd.AddCommand(emr_describeReleaseLabelCmd)
}
