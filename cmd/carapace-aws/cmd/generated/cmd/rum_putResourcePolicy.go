package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Use this operation to assign a resource-based policy to a CloudWatch RUM app monitor to control access to it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_putResourcePolicyCmd).Standalone()

		rum_putResourcePolicyCmd.Flags().String("name", "", "The name of the app monitor that you want to apply this resource-based policy to.")
		rum_putResourcePolicyCmd.Flags().String("policy-document", "", "The JSON to use as the resource policy.")
		rum_putResourcePolicyCmd.Flags().String("policy-revision-id", "", "A string value that you can use to conditionally update your policy.")
		rum_putResourcePolicyCmd.MarkFlagRequired("name")
		rum_putResourcePolicyCmd.MarkFlagRequired("policy-document")
	})
	rumCmd.AddCommand(rum_putResourcePolicyCmd)
}
