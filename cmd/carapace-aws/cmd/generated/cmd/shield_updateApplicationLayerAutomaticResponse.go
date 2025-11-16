package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_updateApplicationLayerAutomaticResponseCmd = &cobra.Command{
	Use:   "update-application-layer-automatic-response",
	Short: "Updates an existing Shield Advanced automatic application layer DDoS mitigation configuration for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_updateApplicationLayerAutomaticResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_updateApplicationLayerAutomaticResponseCmd).Standalone()

		shield_updateApplicationLayerAutomaticResponseCmd.Flags().String("action", "", "Specifies the action setting that Shield Advanced should use in the WAF rules that it creates on behalf of the protected resource in response to DDoS attacks.")
		shield_updateApplicationLayerAutomaticResponseCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the resource.")
		shield_updateApplicationLayerAutomaticResponseCmd.MarkFlagRequired("action")
		shield_updateApplicationLayerAutomaticResponseCmd.MarkFlagRequired("resource-arn")
	})
	shieldCmd.AddCommand(shield_updateApplicationLayerAutomaticResponseCmd)
}
