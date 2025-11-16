package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_listSlackChannelConfigurationsCmd = &cobra.Command{
	Use:   "list-slack-channel-configurations",
	Short: "Lists the Slack channel configurations for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_listSlackChannelConfigurationsCmd).Standalone()

	supportApp_listSlackChannelConfigurationsCmd.Flags().String("next-token", "", "If the results of a search are large, the API only returns a portion of the results and includes a `nextToken` pagination token in the response.")
	supportAppCmd.AddCommand(supportApp_listSlackChannelConfigurationsCmd)
}
