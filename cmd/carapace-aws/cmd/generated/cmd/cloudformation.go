package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformationCmd = &cobra.Command{
	Use:   "cloudformation",
	Short: "CloudFormation\n\nCloudFormation allows you to create and manage Amazon Web Services infrastructure deployments predictably and repeatedly.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformationCmd).Standalone()

	})
	rootCmd.AddCommand(cloudformationCmd)
}
