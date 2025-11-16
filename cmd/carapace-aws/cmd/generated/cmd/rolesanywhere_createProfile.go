package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Creates a *profile*, a list of the roles that Roles Anywhere service is trusted to assume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_createProfileCmd).Standalone()

	rolesanywhere_createProfileCmd.Flags().Bool("accept-role-session-name", false, "Used to determine if a custom role session name will be accepted in a temporary credential request.")
	rolesanywhere_createProfileCmd.Flags().String("duration-seconds", "", "Used to determine how long sessions vended using this profile are valid for.")
	rolesanywhere_createProfileCmd.Flags().Bool("enabled", false, "Specifies whether the profile is enabled.")
	rolesanywhere_createProfileCmd.Flags().String("managed-policy-arns", "", "A list of managed policy ARNs that apply to the vended session credentials.")
	rolesanywhere_createProfileCmd.Flags().String("name", "", "The name of the profile.")
	rolesanywhere_createProfileCmd.Flags().Bool("no-accept-role-session-name", false, "Used to determine if a custom role session name will be accepted in a temporary credential request.")
	rolesanywhere_createProfileCmd.Flags().Bool("no-enabled", false, "Specifies whether the profile is enabled.")
	rolesanywhere_createProfileCmd.Flags().Bool("no-require-instance-properties", false, "Specifies whether instance properties are required in temporary credential requests with this profile.")
	rolesanywhere_createProfileCmd.Flags().Bool("require-instance-properties", false, "Specifies whether instance properties are required in temporary credential requests with this profile.")
	rolesanywhere_createProfileCmd.Flags().String("role-arns", "", "A list of IAM roles that this profile can assume in a temporary credential request.")
	rolesanywhere_createProfileCmd.Flags().String("session-policy", "", "A session policy that applies to the trust boundary of the vended session credentials.")
	rolesanywhere_createProfileCmd.Flags().String("tags", "", "The tags to attach to the profile.")
	rolesanywhere_createProfileCmd.MarkFlagRequired("name")
	rolesanywhere_createProfileCmd.Flag("no-accept-role-session-name").Hidden = true
	rolesanywhere_createProfileCmd.Flag("no-enabled").Hidden = true
	rolesanywhere_createProfileCmd.Flag("no-require-instance-properties").Hidden = true
	rolesanywhere_createProfileCmd.MarkFlagRequired("role-arns")
	rolesanywhereCmd.AddCommand(rolesanywhere_createProfileCmd)
}
