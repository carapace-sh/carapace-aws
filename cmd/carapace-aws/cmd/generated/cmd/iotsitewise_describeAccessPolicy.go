package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeAccessPolicyCmd = &cobra.Command{
	Use:   "describe-access-policy",
	Short: "Describes an access policy, which specifies an identity's access to an IoT SiteWise Monitor portal or project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeAccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeAccessPolicyCmd).Standalone()

		iotsitewise_describeAccessPolicyCmd.Flags().String("access-policy-id", "", "The ID of the access policy.")
		iotsitewise_describeAccessPolicyCmd.MarkFlagRequired("access-policy-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeAccessPolicyCmd)
}
