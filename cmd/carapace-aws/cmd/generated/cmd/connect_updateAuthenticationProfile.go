package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateAuthenticationProfileCmd = &cobra.Command{
	Use:   "update-authentication-profile",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateAuthenticationProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateAuthenticationProfileCmd).Standalone()

		connect_updateAuthenticationProfileCmd.Flags().String("allowed-ips", "", "A list of IP address range strings that are allowed to access the instance.")
		connect_updateAuthenticationProfileCmd.Flags().String("authentication-profile-id", "", "A unique identifier for the authentication profile.")
		connect_updateAuthenticationProfileCmd.Flags().String("blocked-ips", "", "A list of IP address range strings that are blocked from accessing the instance.")
		connect_updateAuthenticationProfileCmd.Flags().String("description", "", "The description for the authentication profile.")
		connect_updateAuthenticationProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateAuthenticationProfileCmd.Flags().String("name", "", "The name for the authentication profile.")
		connect_updateAuthenticationProfileCmd.Flags().Bool("no-session-inactivity-handling-enabled", false, "Determines if automatic logout on user inactivity is enabled.")
		connect_updateAuthenticationProfileCmd.Flags().String("periodic-session-duration", "", "The short lived session duration configuration for users logged in to Amazon Connect, in minutes.")
		connect_updateAuthenticationProfileCmd.Flags().String("session-inactivity-duration", "", "The period, in minutes, before an agent is automatically signed out of the contact center when they go inactive.")
		connect_updateAuthenticationProfileCmd.Flags().Bool("session-inactivity-handling-enabled", false, "Determines if automatic logout on user inactivity is enabled.")
		connect_updateAuthenticationProfileCmd.MarkFlagRequired("authentication-profile-id")
		connect_updateAuthenticationProfileCmd.MarkFlagRequired("instance-id")
		connect_updateAuthenticationProfileCmd.Flag("no-session-inactivity-handling-enabled").Hidden = true
	})
	connectCmd.AddCommand(connect_updateAuthenticationProfileCmd)
}
