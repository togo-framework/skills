// Package skills is the Skills content pack for togo: Skill resource: proficiency-ranked grid by category.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package skills

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("skills", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
