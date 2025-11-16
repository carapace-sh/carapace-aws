package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_getSafetyLeverCmd = &cobra.Command{
	Use:   "get-safety-lever",
	Short: "Gets information about the specified safety lever.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_getSafetyLeverCmd).Standalone()

	fis_getSafetyLeverCmd.Flags().String("id", "", "The ID of the safety lever.")
	fis_getSafetyLeverCmd.MarkFlagRequired("id")
	fisCmd.AddCommand(fis_getSafetyLeverCmd)
}
