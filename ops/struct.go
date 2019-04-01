package ops

type JobStack struct{
	Id string `json:"id,omitempty"`
	Info *Job `json:"info,omitempty"`
}

type JobStackArray []JobStack

type Job struct{
  Title string  `json:"title,omitempty"`
  Short_desc string `json:"shortdesc,omitempty"`
  Coordinates *Coordinates `json:"coordinates,omitempty"`
  Contact string `json:"contact,omitempty"`
  MetaData *JobMeta `json:"meta,omitempty"`
}

type JobMeta struct{
  Added_date  string `json:added_date,omitempty`
  Added_user  string `json:added_user,omitempty`
  Modified_date string `json:modified_date,omitempty`
  Views string `json:views,omitempty`
}

type Coordinates struct{
  Latitude string `json:latitude,omitempty`
  Longtitude string `json:longtitude,omitempty`
}

