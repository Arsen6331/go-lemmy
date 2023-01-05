package types

type Activity struct {
	ID        int32          `json:"id" url:"id,omitempty"`
	Data      any            `json:"data" url:"data,omitempty"`
	Local     bool           `json:"local" url:"local,omitempty"`
	Published LemmyTime      `json:"published" url:"published,omitempty"`
	Updated   LemmyTime      `json:"updated" url:"updated,omitempty"`
	ApID      string         `json:"ap_id" url:"ap_id,omitempty"`
	Sensitive Optional[bool] `json:"sensitive" url:"sensitive,omitempty"`
}
type ActivityForm struct {
	Data      any            `json:"data" url:"data,omitempty"`
	Local     Optional[bool] `json:"local" url:"local,omitempty"`
	Updated   LemmyTime      `json:"updated" url:"updated,omitempty"`
	ApID      string         `json:"ap_id" url:"ap_id,omitempty"`
	Sensitive bool           `json:"sensitive" url:"sensitive,omitempty"`
}