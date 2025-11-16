package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_createAllowListCmd = &cobra.Command{
	Use:   "create-allow-list",
	Short: "Creates and defines the settings for an allow list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_createAllowListCmd).Standalone()

	macie2_createAllowListCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
	macie2_createAllowListCmd.Flags().String("criteria", "", "The criteria that specify the text or text pattern to ignore.")
	macie2_createAllowListCmd.Flags().String("description", "", "A custom description of the allow list.")
	macie2_createAllowListCmd.Flags().String("name", "", "A custom name for the allow list.")
	macie2_createAllowListCmd.Flags().String("tags", "", "A map of key-value pairs that specifies the tags to associate with the allow list.")
	macie2_createAllowListCmd.MarkFlagRequired("client-token")
	macie2_createAllowListCmd.MarkFlagRequired("criteria")
	macie2_createAllowListCmd.MarkFlagRequired("name")
	macie2Cmd.AddCommand(macie2_createAllowListCmd)
}
