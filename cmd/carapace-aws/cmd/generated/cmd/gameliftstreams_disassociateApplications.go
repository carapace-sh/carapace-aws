package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_disassociateApplicationsCmd = &cobra.Command{
	Use:   "disassociate-applications",
	Short: "When you disassociate, or unlink, an application from a stream group, you can no longer stream this application by using that stream group's allocated compute resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_disassociateApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_disassociateApplicationsCmd).Standalone()

		gameliftstreams_disassociateApplicationsCmd.Flags().String("application-identifiers", "", "A set of applications that you want to disassociate from the stream group.")
		gameliftstreams_disassociateApplicationsCmd.Flags().String("identifier", "", "A stream group to disassociate these applications from.")
		gameliftstreams_disassociateApplicationsCmd.MarkFlagRequired("application-identifiers")
		gameliftstreams_disassociateApplicationsCmd.MarkFlagRequired("identifier")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_disassociateApplicationsCmd)
}
