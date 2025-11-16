package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfigdata_startConfigurationSessionCmd = &cobra.Command{
	Use:   "start-configuration-session",
	Short: "Starts a configuration session used to retrieve a deployed configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfigdata_startConfigurationSessionCmd).Standalone()

	appconfigdata_startConfigurationSessionCmd.Flags().String("application-identifier", "", "The application ID or the application name.")
	appconfigdata_startConfigurationSessionCmd.Flags().String("configuration-profile-identifier", "", "The configuration profile ID or the configuration profile name.")
	appconfigdata_startConfigurationSessionCmd.Flags().String("environment-identifier", "", "The environment ID or the environment name.")
	appconfigdata_startConfigurationSessionCmd.Flags().String("required-minimum-poll-interval-in-seconds", "", "Sets a constraint on a session.")
	appconfigdata_startConfigurationSessionCmd.MarkFlagRequired("application-identifier")
	appconfigdata_startConfigurationSessionCmd.MarkFlagRequired("configuration-profile-identifier")
	appconfigdata_startConfigurationSessionCmd.MarkFlagRequired("environment-identifier")
	appconfigdataCmd.AddCommand(appconfigdata_startConfigurationSessionCmd)
}
