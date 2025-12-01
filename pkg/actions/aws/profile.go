package aws

import (
	"github.com/carapace-sh/carapace"
	"gopkg.in/ini.v1"
	"strings"
)

// ActionProfiles completes configuration profile names
//
//	someprofile (eu-central-1)
//	anotherprofile (us-east-1)
func ActionProfiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		profiles := []string{}

		// TODO support windows
		if path, err := c.Abs("~/.aws/config"); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if cfg, err := ini.Load(path); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				for _, section := range cfg.Sections() {
					if after, ok := strings.CutPrefix(section.Name(), "profile "); ok {
						profiles = append(profiles, after)
						if key, err := section.GetKey("region"); err != nil {
							profiles = append(profiles, "")
						} else {
							profiles = append(profiles, key.String())
						}
					}
				}
				if len(profiles) == 0 {
					profiles = append(profiles, "default", "")
				}
				return carapace.ActionValuesDescribed(profiles...)
			}
		}
	})
}
