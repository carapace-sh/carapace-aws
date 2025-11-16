package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Sets the resource policy to grant one or more Amazon Web Services services and accounts permissions to access X-Ray.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_putResourcePolicyCmd).Standalone()

		xray_putResourcePolicyCmd.Flags().Bool("bypass-policy-lockout-check", false, "A flag to indicate whether to bypass the resource policy lockout safety check.")
		xray_putResourcePolicyCmd.Flags().Bool("no-bypass-policy-lockout-check", false, "A flag to indicate whether to bypass the resource policy lockout safety check.")
		xray_putResourcePolicyCmd.Flags().String("policy-document", "", "The resource policy document, which can be up to 5kb in size.")
		xray_putResourcePolicyCmd.Flags().String("policy-name", "", "The name of the resource policy.")
		xray_putResourcePolicyCmd.Flags().String("policy-revision-id", "", "Specifies a specific policy revision, to ensure an atomic create operation.")
		xray_putResourcePolicyCmd.Flag("no-bypass-policy-lockout-check").Hidden = true
		xray_putResourcePolicyCmd.MarkFlagRequired("policy-document")
		xray_putResourcePolicyCmd.MarkFlagRequired("policy-name")
	})
	xrayCmd.AddCommand(xray_putResourcePolicyCmd)
}
