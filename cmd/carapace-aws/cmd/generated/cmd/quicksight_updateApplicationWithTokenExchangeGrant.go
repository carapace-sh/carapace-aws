package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateApplicationWithTokenExchangeGrantCmd = &cobra.Command{
	Use:   "update-application-with-token-exchange-grant",
	Short: "Updates an Quick Suite application with a token exchange grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateApplicationWithTokenExchangeGrantCmd).Standalone()

	quicksight_updateApplicationWithTokenExchangeGrantCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account to be updated with a token exchange grant.")
	quicksight_updateApplicationWithTokenExchangeGrantCmd.Flags().String("namespace", "", "The namespace of the Quick Suite application.")
	quicksight_updateApplicationWithTokenExchangeGrantCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateApplicationWithTokenExchangeGrantCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_updateApplicationWithTokenExchangeGrantCmd)
}
