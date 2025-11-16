package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeEntityCmd = &cobra.Command{
	Use:   "describe-entity",
	Short: "Returns basic details about an entity in WorkMail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeEntityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_describeEntityCmd).Standalone()

		workmail_describeEntityCmd.Flags().String("email", "", "The email under which the entity exists.")
		workmail_describeEntityCmd.Flags().String("organization-id", "", "The identifier for the organization under which the entity exists.")
		workmail_describeEntityCmd.MarkFlagRequired("email")
		workmail_describeEntityCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_describeEntityCmd)
}
