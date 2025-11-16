package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var freetier_getAccountActivityCmd = &cobra.Command{
	Use:   "get-account-activity",
	Short: "Returns a specific activity record that is available to the customer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freetier_getAccountActivityCmd).Standalone()

	freetier_getAccountActivityCmd.Flags().String("activity-id", "", "A unique identifier that identifies the activity.")
	freetier_getAccountActivityCmd.Flags().String("language-code", "", "The language code used to return translated title and description fields.")
	freetier_getAccountActivityCmd.MarkFlagRequired("activity-id")
	freetierCmd.AddCommand(freetier_getAccountActivityCmd)
}
