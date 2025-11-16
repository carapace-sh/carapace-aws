package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_deleteCustomDataIdentifierCmd = &cobra.Command{
	Use:   "delete-custom-data-identifier",
	Short: "Soft deletes a custom data identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_deleteCustomDataIdentifierCmd).Standalone()

	macie2_deleteCustomDataIdentifierCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_deleteCustomDataIdentifierCmd.MarkFlagRequired("id")
	macie2Cmd.AddCommand(macie2_deleteCustomDataIdentifierCmd)
}
