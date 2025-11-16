package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getTestGridSessionCmd = &cobra.Command{
	Use:   "get-test-grid-session",
	Short: "A session is an instance of a browser created through a `RemoteWebDriver` with the URL from [CreateTestGridUrlResult$url]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getTestGridSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getTestGridSessionCmd).Standalone()

		devicefarm_getTestGridSessionCmd.Flags().String("project-arn", "", "The ARN for the project that this session belongs to.")
		devicefarm_getTestGridSessionCmd.Flags().String("session-arn", "", "An ARN that uniquely identifies a [TestGridSession]().")
		devicefarm_getTestGridSessionCmd.Flags().String("session-id", "", "An ID associated with this session.")
	})
	devicefarmCmd.AddCommand(devicefarm_getTestGridSessionCmd)
}
