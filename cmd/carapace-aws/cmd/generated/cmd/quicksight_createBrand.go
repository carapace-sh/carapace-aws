package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createBrandCmd = &cobra.Command{
	Use:   "create-brand",
	Short: "Creates an Quick Sight brand.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createBrandCmd).Standalone()

	quicksight_createBrandCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand.")
	quicksight_createBrandCmd.Flags().String("brand-definition", "", "The definition of the brand.")
	quicksight_createBrandCmd.Flags().String("brand-id", "", "The ID of the Quick Suite brand.")
	quicksight_createBrandCmd.Flags().String("tags", "", "A map of the key-value pairs that are assigned to the brand.")
	quicksight_createBrandCmd.MarkFlagRequired("aws-account-id")
	quicksight_createBrandCmd.MarkFlagRequired("brand-id")
	quicksightCmd.AddCommand(quicksight_createBrandCmd)
}
