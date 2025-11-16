package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listOnPremisesInstancesCmd = &cobra.Command{
	Use:   "list-on-premises-instances",
	Short: "Gets a list of names for one or more on-premises instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listOnPremisesInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listOnPremisesInstancesCmd).Standalone()

		codedeploy_listOnPremisesInstancesCmd.Flags().String("next-token", "", "An identifier returned from the previous list on-premises instances call.")
		codedeploy_listOnPremisesInstancesCmd.Flags().String("registration-status", "", "The registration status of the on-premises instances:")
		codedeploy_listOnPremisesInstancesCmd.Flags().String("tag-filters", "", "The on-premises instance tags that are used to restrict the on-premises instance names returned.")
	})
	codedeployCmd.AddCommand(codedeploy_listOnPremisesInstancesCmd)
}
