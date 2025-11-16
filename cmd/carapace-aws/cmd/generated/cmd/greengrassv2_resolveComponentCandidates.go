package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_resolveComponentCandidatesCmd = &cobra.Command{
	Use:   "resolve-component-candidates",
	Short: "Retrieves a list of components that meet the component, version, and platform requirements of a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_resolveComponentCandidatesCmd).Standalone()

	greengrassv2_resolveComponentCandidatesCmd.Flags().String("component-candidates", "", "The list of components to resolve.")
	greengrassv2_resolveComponentCandidatesCmd.Flags().String("platform", "", "The platform to use to resolve compatible components.")
	greengrassv2Cmd.AddCommand(greengrassv2_resolveComponentCandidatesCmd)
}
