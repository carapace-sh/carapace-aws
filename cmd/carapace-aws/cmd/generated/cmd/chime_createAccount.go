package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_createAccountCmd = &cobra.Command{
	Use:   "create-account",
	Short: "Creates an Amazon Chime account under the administrator's AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_createAccountCmd).Standalone()

	chime_createAccountCmd.Flags().String("name", "", "The name of the Amazon Chime account.")
	chime_createAccountCmd.MarkFlagRequired("name")
	chimeCmd.AddCommand(chime_createAccountCmd)
}
