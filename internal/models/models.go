package models

type User struct {
	Username     string                 `json:"-"`
	FullName     string                 `json:"full_name"`
	Email        string                 `json:"email"`
	Roles        []string               `json:"roles"`
	Password     *string                `json:"password,omitempty"`
	PasswordHash *string                `json:"password_hash,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Enabled      bool                   `json:"enabled"`
}

type Role struct {
	Name         string                 `json:"-"`
	Applications []Application          `json:"applications,omitempty"`
	Global       map[string]interface{} `json:"global,omitempty"`
	Cluster      []string               `json:"cluster,omitempty"`
	Indices      []IndexPerms           `json:"indices,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	RusAs        []string               `json:"run_as,omitempty"`
}

type IndexPerms struct {
	FieldSecurity *FieldSecurity `json:"field_security,omitempty"`
	Names         []string       `json:"names"`
	Privileges    []string       `json:"privileges"`
	Query         *string        `json:"query,omitempty"`
}

type FieldSecurity struct {
	Grant  []string `json:"grant,omitempty"`
	Except []string `json:"except,omitempty"`
}

type Application struct {
	Name       string   `json:"application"`
	Privileges []string `json:"privileges,omitempty"`
	Resources  []string `json:"resources"`
}

type IndexTemplate struct {
	Name          string                 `json:"-"`
	Create        bool                   `json:"-"`
	Timeout       string                 `json:"-"`
	ComposedOf    []string               `json:"composed_of"`
	DataStream    map[string]interface{} `json:"data_stream,omitempty"`
	IndexPatterns []string               `json:"index_patterns"`
	Meta          map[string]interface{} `json:"_meta,omitempty"`
	Priority      *int                   `json:"priority,omitempty"`
	Template      *Template              `json:"template,omitempty"`
	Version       *int                   `json:"version,omitempty"`
}

type Template struct {
	Aliases  map[string]TemplateAlias `json:"aliases,omitempty"`
	Mappings map[string]interface{}   `json:"mappings,omitempty"`
	Settings map[string]interface{}   `json:"settings,omitempty"`
}

type TemplateAlias struct {
	Name          string                 `json:"-"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	IndexRouting  string                 `json:"index_routing,omitempty"`
	IsHidden      bool                   `json:"is_hidden,omitempty"`
	IsWriteIndex  bool                   `json:"is_write_index,omitempty"`
	Routing       string                 `json:"routing,omitempty"`
	SearchRouting string                 `json:"search_routing,omitempty"`
}

type IndexTemplatesResponse struct {
	IndexTemplates []IndexTemplateResponse `json:"index_templates"`
}

type IndexTemplateResponse struct {
	Name          string        `json:"name"`
	IndexTemplate IndexTemplate `json:"index_template"`
}

type PolicyDefinition struct {
	Policy   Policy `json:"policy"`
	Modified string `json:"modified_date"`
}

type Policy struct {
	Name     string                 `json:"-"`
	Metadata map[string]interface{} `json:"_meta,omitempty"`
	Phases   map[string]Phase       `json:"phases"`
}

type Phase struct {
	MinAge  string            `json:"min_age,omitempty"`
	Actions map[string]Action `json:"actions"`
}

type Action map[string]interface{}

type SnapshotRepository struct {
	Name     string                 `json:"-"`
	Type     string                 `json:"type"`
	Settings map[string]interface{} `json:"settings"`
	Verify   bool                   `json:"verify"`
}

type SnapshotPolicy struct {
	Id         string                `json:"-"`
	Config     *SnapshotPolicyConfig `json:"config,omitempty"`
	Name       string                `json:"name"`
	Repository string                `json:"repository"`
	Retention  *SnapshotRetention    `json:"retention,omitempty"`
	Schedule   string                `json:"schedule"`
}

type SnapshotRetention struct {
	ExpireAfter *string `json:"expire_after,omitempty"`
	MaxCount    *int    `json:"max_count,omitempty"`
	MinCount    *int    `json:"min_count,omitempty"`
}

type SnapshotPolicyConfig struct {
	ExpandWildcards    *string                `json:"expand_wildcards,omitempty"`
	IgnoreUnavailable  *bool                  `json:"ignore_unavailable,omitempty"`
	IncludeGlobalState *bool                  `json:"include_global_state,omitempty"`
	Indices            []string               `json:"indices,omitempty"`
	FeatureStates      []string               `json:"feature_states,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	Partial            *bool                  `json:"partial,omitempty"`
}
