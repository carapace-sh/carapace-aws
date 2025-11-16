package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getOnPremisesInstanceCmd = &cobra.Command{
	Use:   "get-on-premises-instance",
	Short: "Gets information about an on-premises instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getOnPremisesInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_getOnPremisesInstanceCmd).Standalone()

		codedeploy_getOnPremisesInstanceCmd.Flags().String("instance-name", "", "The name of the on-premises instance about which to get information.")
		codedeploy_getOnPremisesInstanceCmd.MarkFlagRequired("instance-name")
	})
	codedeployCmd.AddCommand(codedeploy_getOnPremisesInstanceCmd)
}
