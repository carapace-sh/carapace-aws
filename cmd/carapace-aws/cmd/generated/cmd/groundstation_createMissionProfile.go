package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_createMissionProfileCmd = &cobra.Command{
	Use:   "create-mission-profile",
	Short: "Creates a mission profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_createMissionProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_createMissionProfileCmd).Standalone()

		groundstation_createMissionProfileCmd.Flags().String("contact-post-pass-duration-seconds", "", "Amount of time after a contact ends that you’d like to receive a Ground Station Contact State Change event indicating the pass has finished.")
		groundstation_createMissionProfileCmd.Flags().String("contact-pre-pass-duration-seconds", "", "Amount of time prior to contact start you’d like to receive a Ground Station Contact State Change event indicating an upcoming pass.")
		groundstation_createMissionProfileCmd.Flags().String("dataflow-edges", "", "A list of lists of ARNs.")
		groundstation_createMissionProfileCmd.Flags().String("minimum-viable-contact-duration-seconds", "", "Smallest amount of time in seconds that you’d like to see for an available contact.")
		groundstation_createMissionProfileCmd.Flags().String("name", "", "Name of a mission profile.")
		groundstation_createMissionProfileCmd.Flags().String("streams-kms-key", "", "KMS key to use for encrypting streams.")
		groundstation_createMissionProfileCmd.Flags().String("streams-kms-role", "", "Role to use for encrypting streams with KMS key.")
		groundstation_createMissionProfileCmd.Flags().String("tags", "", "Tags assigned to a mission profile.")
		groundstation_createMissionProfileCmd.Flags().String("tracking-config-arn", "", "ARN of a tracking `Config`.")
		groundstation_createMissionProfileCmd.MarkFlagRequired("dataflow-edges")
		groundstation_createMissionProfileCmd.MarkFlagRequired("minimum-viable-contact-duration-seconds")
		groundstation_createMissionProfileCmd.MarkFlagRequired("name")
		groundstation_createMissionProfileCmd.MarkFlagRequired("tracking-config-arn")
	})
	groundstationCmd.AddCommand(groundstation_createMissionProfileCmd)
}
