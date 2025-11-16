package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeKeyRegistrationCmd = &cobra.Command{
	Use:   "describe-key-registration",
	Short: "Describes all customer managed key registrations in a Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeKeyRegistrationCmd).Standalone()

	quicksight_describeKeyRegistrationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the customer managed key registration that you want to describe.")
	quicksight_describeKeyRegistrationCmd.Flags().Bool("default-key-only", false, "Determines whether the request returns the default key only.")
	quicksight_describeKeyRegistrationCmd.Flags().Bool("no-default-key-only", false, "Determines whether the request returns the default key only.")
	quicksight_describeKeyRegistrationCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeKeyRegistrationCmd.Flag("no-default-key-only").Hidden = true
	quicksightCmd.AddCommand(quicksight_describeKeyRegistrationCmd)
}
