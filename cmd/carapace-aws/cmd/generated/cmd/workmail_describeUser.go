package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeUserCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "Provides information regarding the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeUserCmd).Standalone()

	workmail_describeUserCmd.Flags().String("organization-id", "", "The identifier for the organization under which the user exists.")
	workmail_describeUserCmd.Flags().String("user-id", "", "The identifier for the user to be described.")
	workmail_describeUserCmd.MarkFlagRequired("organization-id")
	workmail_describeUserCmd.MarkFlagRequired("user-id")
	workmailCmd.AddCommand(workmail_describeUserCmd)
}
