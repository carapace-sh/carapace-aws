package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeEventSourcesConfigCmd = &cobra.Command{
	Use:   "describe-event-sources-config",
	Short: "Returns the integration status of services that are integrated with DevOps Guru as Consumer via EventBridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeEventSourcesConfigCmd).Standalone()

	devopsGuruCmd.AddCommand(devopsGuru_describeEventSourcesConfigCmd)
}
