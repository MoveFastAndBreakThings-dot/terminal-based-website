package tests

import (
	"strings"
	"testing"

	"ssh-portfolio/content"
)

// ── Profile ───────────────────────────────────────────────────────────────────

func TestProfile_NameNotEmpty(t *testing.T) {
	if strings.TrimSpace(content.MyProfile.Name) == "" {
		t.Fatal("Profile.Name is empty")
	}
}

func TestProfile_RoleNotEmpty(t *testing.T) {
	if strings.TrimSpace(content.MyProfile.Role) == "" {
		t.Fatal("Profile.Role is empty")
	}
}

func TestProfile_BioHasParagraphs(t *testing.T) {
	if len(content.MyProfile.Bio) == 0 {
		t.Fatal("Profile.Bio has no paragraphs")
	}
	for i, para := range content.MyProfile.Bio {
		if strings.TrimSpace(para) == "" {
			t.Errorf("Profile.Bio[%d] is blank", i)
		}
	}
}

// ── Jobs ──────────────────────────────────────────────────────────────────────

func TestJobs_NotEmpty(t *testing.T) {
	if len(content.Jobs) == 0 {
		t.Fatal("Jobs slice is empty")
	}
}

func TestJobs_RequiredFields(t *testing.T) {
	for i, job := range content.Jobs {
		if strings.TrimSpace(job.Title) == "" {
			t.Errorf("Jobs[%d].Title is empty", i)
		}
		if strings.TrimSpace(job.Company) == "" {
			t.Errorf("Jobs[%d].Company is empty", i)
		}
		if strings.TrimSpace(job.Period) == "" {
			t.Errorf("Jobs[%d].Period is empty", i)
		}
	}
}

func TestJobs_BulletsNotEmpty(t *testing.T) {
	for i, job := range content.Jobs {
		if len(job.Bullets) == 0 {
			t.Errorf("Jobs[%d] (%s) has no bullets", i, job.Title)
		}
		for j, b := range job.Bullets {
			if strings.TrimSpace(b) == "" {
				t.Errorf("Jobs[%d].Bullets[%d] is blank", i, j)
			}
		}
	}
}

func TestJobs_TagsNotEmpty(t *testing.T) {
	for i, job := range content.Jobs {
		if len(job.Tags) == 0 {
			t.Errorf("Jobs[%d] (%s) has no tags", i, job.Title)
		}
		for j, tag := range job.Tags {
			if strings.TrimSpace(tag) == "" {
				t.Errorf("Jobs[%d].Tags[%d] is blank", i, j)
			}
		}
	}
}

// ── Projects ──────────────────────────────────────────────────────────────────

func TestProjects_NotEmpty(t *testing.T) {
	if len(content.Projects) == 0 {
		t.Fatal("Projects slice is empty")
	}
}

func TestProjects_RequiredFields(t *testing.T) {
	for i, proj := range content.Projects {
		if strings.TrimSpace(proj.Name) == "" {
			t.Errorf("Projects[%d].Name is empty", i)
		}
		if strings.TrimSpace(proj.Description) == "" {
			t.Errorf("Projects[%d].Description is empty", i)
		}
	}
}

func TestProjects_TagsNotEmpty(t *testing.T) {
	for i, proj := range content.Projects {
		if len(proj.Tags) == 0 {
			t.Errorf("Projects[%d] (%s) has no tags", i, proj.Name)
		}
		for j, tag := range proj.Tags {
			if strings.TrimSpace(tag) == "" {
				t.Errorf("Projects[%d].Tags[%d] is blank", i, j)
			}
		}
	}
}

func TestProjects_URLsWellFormed(t *testing.T) {
	for i, proj := range content.Projects {
		if proj.URL == "" {
			continue
		}
		if !strings.HasPrefix(proj.URL, "http://") && !strings.HasPrefix(proj.URL, "https://") {
			t.Errorf("Projects[%d].URL %q not http/https", i, proj.URL)
		}
	}
}

// ── SkillGroups ───────────────────────────────────────────────────────────────

func TestSkillGroups_NotEmpty(t *testing.T) {
	if len(content.SkillGroups) == 0 {
		t.Fatal("SkillGroups slice is empty")
	}
}

func TestSkillGroups_CategoryNotEmpty(t *testing.T) {
	for i, sg := range content.SkillGroups {
		if strings.TrimSpace(sg.Category) == "" {
			t.Errorf("SkillGroups[%d].Category is empty", i)
		}
	}
}

func TestSkillGroups_EachHasItems(t *testing.T) {
	for i, sg := range content.SkillGroups {
		if len(sg.Items) == 0 {
			t.Errorf("SkillGroups[%d] (%s) has no items", i, sg.Category)
		}
		for j, item := range sg.Items {
			if strings.TrimSpace(item) == "" {
				t.Errorf("SkillGroups[%d].Items[%d] is blank", i, j)
			}
		}
	}
}

func TestSkillGroups_NoDuplicateCategories(t *testing.T) {
	seen := map[string]bool{}
	for _, sg := range content.SkillGroups {
		if seen[sg.Category] {
			t.Errorf("duplicate skill category: %q", sg.Category)
		}
		seen[sg.Category] = true
	}
}

// ── Links ─────────────────────────────────────────────────────────────────────

func TestLinks_NotEmpty(t *testing.T) {
	if len(content.Links) == 0 {
		t.Fatal("Links slice is empty")
	}
}

func TestLinks_RequiredFields(t *testing.T) {
	for i, lk := range content.Links {
		if strings.TrimSpace(lk.Label) == "" {
			t.Errorf("Links[%d].Label is empty", i)
		}
		if strings.TrimSpace(lk.URL) == "" {
			t.Errorf("Links[%d].URL is empty", i)
		}
	}
}

func TestLinks_URLsWellFormed(t *testing.T) {
	for i, lk := range content.Links {
		if !strings.HasPrefix(lk.URL, "http://") &&
			!strings.HasPrefix(lk.URL, "https://") &&
			!strings.HasPrefix(lk.URL, "mailto:") {
			t.Errorf("Links[%d] (%s) URL %q not http/https/mailto", i, lk.Label, lk.URL)
		}
	}
}

func TestLinks_NoDuplicateLabels(t *testing.T) {
	seen := map[string]bool{}
	for _, lk := range content.Links {
		if seen[lk.Label] {
			t.Errorf("duplicate link label: %q", lk.Label)
		}
		seen[lk.Label] = true
	}
}

// ── Extracurriculars ──────────────────────────────────────────────────────────

func TestExtracurriculars_NotEmpty(t *testing.T) {
	if len(content.Extracurriculars) == 0 {
		t.Fatal("Extracurriculars slice is empty")
	}
}

func TestExtracurriculars_NoBlanks(t *testing.T) {
	for i, ec := range content.Extracurriculars {
		if strings.TrimSpace(ec) == "" {
			t.Errorf("Extracurriculars[%d] is blank", i)
		}
	}
}
