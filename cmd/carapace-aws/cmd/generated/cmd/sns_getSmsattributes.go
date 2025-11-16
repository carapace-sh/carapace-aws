package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_getSmsattributesCmd = &cobra.Command{
	Use:   "get-smsattributes",
	Short: "Returns the settings for sending SMS messages from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_getSmsattributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_getSmsattributesCmd).Standalone()

		sns_getSmsattributesCmd.Flags().String("attributes", "", "A list of the individual attribute names, such as `MonthlySpendLimit`, for which you want values.")
	})
	snsCmd.AddCommand(sns_getSmsattributesCmd)
}
