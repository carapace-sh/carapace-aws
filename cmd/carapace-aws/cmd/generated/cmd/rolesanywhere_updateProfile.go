package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_updateProfileCmd = &cobra.Command{
	Use:   "update-profile",
	Short: "Updates a *profile*, a list of the roles that IAM Roles Anywhere service is trusted to assume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_updateProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_updateProfileCmd).Standalone()

		rolesanywhere_updateProfileCmd.Flags().Bool("accept-role-session-name", false, "Used to determine if a custom role session name will be accepted in a temporary credential request.")
		rolesanywhere_updateProfileCmd.Flags().String("duration-seconds", "", "Used to determine how long sessions vended using this profile are valid for.")
		rolesanywhere_updateProfileCmd.Flags().String("managed-policy-arns", "", "A list of managed policy ARNs that apply to the vended session credentials.")
		rolesanywhere_updateProfileCmd.Flags().String("name", "", "The name of the profile.")
		rolesanywhere_updateProfileCmd.Flags().Bool("no-accept-role-session-name", false, "Used to determine if a custom role session name will be accepted in a temporary credential request.")
		rolesanywhere_updateProfileCmd.Flags().String("profile-id", "", "The unique identifier of the profile.")
		rolesanywhere_updateProfileCmd.Flags().String("role-arns", "", "A list of IAM roles that this profile can assume in a temporary credential request.")
		rolesanywhere_updateProfileCmd.Flags().String("session-policy", "", "A session policy that applies to the trust boundary of the vended session credentials.")
		rolesanywhere_updateProfileCmd.Flag("no-accept-role-session-name").Hidden = true
		rolesanywhere_updateProfileCmd.MarkFlagRequired("profile-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_updateProfileCmd)
}
