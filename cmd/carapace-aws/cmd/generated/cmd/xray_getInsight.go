package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getInsightCmd = &cobra.Command{
	Use:   "get-insight",
	Short: "Retrieves the summary information of an insight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getInsightCmd).Standalone()

	xray_getInsightCmd.Flags().String("insight-id", "", "The insight's unique identifier.")
	xray_getInsightCmd.MarkFlagRequired("insight-id")
	xrayCmd.AddCommand(xray_getInsightCmd)
}
