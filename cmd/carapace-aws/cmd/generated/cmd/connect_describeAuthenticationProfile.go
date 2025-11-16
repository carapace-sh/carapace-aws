package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeAuthenticationProfileCmd = &cobra.Command{
	Use:   "describe-authentication-profile",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeAuthenticationProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeAuthenticationProfileCmd).Standalone()

		connect_describeAuthenticationProfileCmd.Flags().String("authentication-profile-id", "", "A unique identifier for the authentication profile.")
		connect_describeAuthenticationProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeAuthenticationProfileCmd.MarkFlagRequired("authentication-profile-id")
		connect_describeAuthenticationProfileCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_describeAuthenticationProfileCmd)
}
