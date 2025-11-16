package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_deregisterOnPremisesInstanceCmd = &cobra.Command{
	Use:   "deregister-on-premises-instance",
	Short: "Deregisters an on-premises instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_deregisterOnPremisesInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_deregisterOnPremisesInstanceCmd).Standalone()

		codedeploy_deregisterOnPremisesInstanceCmd.Flags().String("instance-name", "", "The name of the on-premises instance to deregister.")
		codedeploy_deregisterOnPremisesInstanceCmd.MarkFlagRequired("instance-name")
	})
	codedeployCmd.AddCommand(codedeploy_deregisterOnPremisesInstanceCmd)
}
