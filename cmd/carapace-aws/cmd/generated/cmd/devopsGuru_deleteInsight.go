package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_deleteInsightCmd = &cobra.Command{
	Use:   "delete-insight",
	Short: "Deletes the insight along with the associated anomalies, events and recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_deleteInsightCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_deleteInsightCmd).Standalone()

		devopsGuru_deleteInsightCmd.Flags().String("id", "", "The ID of the insight.")
		devopsGuru_deleteInsightCmd.MarkFlagRequired("id")
	})
	devopsGuruCmd.AddCommand(devopsGuru_deleteInsightCmd)
}
