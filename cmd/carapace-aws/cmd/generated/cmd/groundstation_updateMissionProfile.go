package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_updateMissionProfileCmd = &cobra.Command{
	Use:   "update-mission-profile",
	Short: "Updates a mission profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_updateMissionProfileCmd).Standalone()

	groundstation_updateMissionProfileCmd.Flags().String("contact-post-pass-duration-seconds", "", "Amount of time after a contact ends that you’d like to receive a Ground Station Contact State Change event indicating the pass has finished.")
	groundstation_updateMissionProfileCmd.Flags().String("contact-pre-pass-duration-seconds", "", "Amount of time after a contact ends that you’d like to receive a Ground Station Contact State Change event indicating the pass has finished.")
	groundstation_updateMissionProfileCmd.Flags().String("dataflow-edges", "", "A list of lists of ARNs.")
	groundstation_updateMissionProfileCmd.Flags().String("minimum-viable-contact-duration-seconds", "", "Smallest amount of time in seconds that you’d like to see for an available contact.")
	groundstation_updateMissionProfileCmd.Flags().String("mission-profile-id", "", "UUID of a mission profile.")
	groundstation_updateMissionProfileCmd.Flags().String("name", "", "Name of a mission profile.")
	groundstation_updateMissionProfileCmd.Flags().String("streams-kms-key", "", "KMS key to use for encrypting streams.")
	groundstation_updateMissionProfileCmd.Flags().String("streams-kms-role", "", "Role to use for encrypting streams with KMS key.")
	groundstation_updateMissionProfileCmd.Flags().String("tracking-config-arn", "", "ARN of a tracking `Config`.")
	groundstation_updateMissionProfileCmd.MarkFlagRequired("mission-profile-id")
	groundstationCmd.AddCommand(groundstation_updateMissionProfileCmd)
}
