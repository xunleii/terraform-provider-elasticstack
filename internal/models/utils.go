package models

import (
	"encoding/json"
)

func (t Template) ToMap() (map[string]interface{}, error) {
	flatmap := make(map[string]interface{})

	if t.Aliases != nil {
		aliases := make([]interface{}, 0, len(t.Aliases))
		for k, v := range t.Aliases {
			alias := make(map[string]interface{})
			alias["name"] = k
			alias["index_routing"] = v.IndexRouting
			alias["is_hidden"] = v.IsHidden
			alias["is_write_index"] = v.IsWriteIndex
			alias["routing"] = v.Routing
			alias["search_routing"] = v.SearchRouting

			if v.Filter != nil {
				if buf, err := json.Marshal(v.Filter); err != nil {
					return nil, err
				} else {
					alias["filter"] = string(buf)
				}
			}

			aliases = append(aliases, alias)
		}
		flatmap["aliases"] = aliases
	}

	if t.Mappings != nil {
		if buf, err := json.Marshal(t.Mappings); err != nil {
			return nil, err
		} else {
			flatmap["mappings"] = string(buf)
		}
	}

	if t.Settings != nil {
		if buf, err := json.Marshal(t.Settings); err != nil {
			return nil, err
		} else {
			flatmap["settings"] = string(buf)
		}
	}

	return flatmap, nil
}
