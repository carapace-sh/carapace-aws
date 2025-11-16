package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_stopApplicationCmd = &cobra.Command{
	Use:   "stop-application",
	Short: "Request is an operation to stop an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_stopApplicationCmd).Standalone()

	ssmSap_stopApplicationCmd.Flags().String("application-id", "", "The ID of the application.")
	ssmSap_stopApplicationCmd.Flags().Bool("include-ec2-instance-shutdown", false, "Boolean.")
	ssmSap_stopApplicationCmd.Flags().Bool("no-include-ec2-instance-shutdown", false, "Boolean.")
	ssmSap_stopApplicationCmd.Flags().String("stop-connected-entity", "", "Specify the `ConnectedEntityType`.")
	ssmSap_stopApplicationCmd.MarkFlagRequired("application-id")
	ssmSap_stopApplicationCmd.Flag("no-include-ec2-instance-shutdown").Hidden = true
	ssmSapCmd.AddCommand(ssmSap_stopApplicationCmd)
}
