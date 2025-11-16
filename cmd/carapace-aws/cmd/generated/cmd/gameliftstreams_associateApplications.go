package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_associateApplicationsCmd = &cobra.Command{
	Use:   "associate-applications",
	Short: "When you associate, or link, an application with a stream group, then Amazon GameLift Streams can launch the application using the stream group's allocated compute resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_associateApplicationsCmd).Standalone()

	gameliftstreams_associateApplicationsCmd.Flags().String("application-identifiers", "", "A set of applications to associate with the stream group.")
	gameliftstreams_associateApplicationsCmd.Flags().String("identifier", "", "A stream group to associate to the applications.")
	gameliftstreams_associateApplicationsCmd.MarkFlagRequired("application-identifiers")
	gameliftstreams_associateApplicationsCmd.MarkFlagRequired("identifier")
	gameliftstreamsCmd.AddCommand(gameliftstreams_associateApplicationsCmd)
}
